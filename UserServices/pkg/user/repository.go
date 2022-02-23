package user

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	ent "github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/entities"
	erro "github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/errors"
	"github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/utils"
)

const collection = "Users" // Equivalent to "table"

type repo struct {
	db     *mongo.Database
	logger log.Logger
}

func NewRepo(logger log.Logger, db *mongo.Database) *repo {
	return &repo{
		db:     db,
		logger: logger,
	}
}

// Get user from database by id or email
func (repo *repo) GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error) {
	logger := log.With(repo.logger, "Repository method", "GetUser")
	var resp ent.GetUserResp
	collection := repo.db.Collection(collection)

	bsonDocumentFilter := utils.GetCorrectFilterId(r.Id)
	err := collection.FindOne(ctx, bsonDocumentFilter).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			level.Error(logger).Log("error", "User not found")
			return ent.GetUserResp{}, erro.NewErrUserNotFound()
		}
		level.Error(logger).Log("error", err)
		return ent.GetUserResp{}, err
	}
	return resp, nil
}

// Create new user in the database
func (repo *repo) CreateUser(ctx context.Context, user ent.User) (ent.CreateUserResp, error) {
	logger := log.With(repo.logger, "Repository method", "CreateUser")
	collection := repo.db.Collection(collection)
	bsonDocument, _ := utils.StructtoBson(user)

	// check (by email) if user already exists in the database
	if _, err := repo.GetUser(ctx, ent.GetUserReq{Id: user.Email}); err == nil {
		level.Error(logger).Log("error", "User already exists")
		return ent.CreateUserResp{}, erro.NewErrAlreadyExists()
	}

	_, err := collection.InsertOne(ctx, bsonDocument)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.CreateUserResp{}, err
	}
	return ent.CreateUserResp{
		Id:      user.Id,
		Created: user.Created,
	}, nil
}

// Delete user from database by id or email
func (repo *repo) DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error) {
	logger := log.With(repo.logger, "Repository method", "Delete User")
	collection := repo.db.Collection(collection)
	var resp ent.DeleteUserResp

	bsonDocumentFilter := utils.GetCorrectFilterId(r.Id)
	err := collection.FindOneAndDelete(ctx, bsonDocumentFilter).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			level.Error(logger).Log("error", "User not found")
			return ent.DeleteUserResp{}, erro.NewErrUserNotFound()
		}
		level.Error(logger).Log("error", err)
		return ent.DeleteUserResp{}, err
	}
	return ent.DeleteUserResp{
		Deleted: utils.TimeNow(),
	}, nil
}

// Create new user in the database
func (repo *repo) UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error) {
	logger := log.With(repo.logger, "Repository method", "UpdateUser")
	collection := repo.db.Collection(collection)
	var resp ent.UpdateUserResp
	update, _ := utils.StructtoBson(r)

	err := collection.FindOneAndUpdate(ctx, bson.D{{"email", r.Email}}, bson.D{{"$set", update}}).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			level.Error(logger).Log("error", "User not found")
			return ent.UpdateUserResp{}, erro.NewErrUserNotFound()
		}
		level.Error(logger).Log("error", err)
		return ent.UpdateUserResp{}, err
	}

	return ent.UpdateUserResp{
		Updated: utils.TimeNow(),
	}, nil
}

func (repo *repo) AuthenticateUser(ctx context.Context, r ent.AuthenticateReq) (string, error) {
	logger := log.With(repo.logger, "Repository method", "AuthenticateUser")
	var resp struct {
		PwdHsh string `bson:"pwdhsh"`
	}
	collection := repo.db.Collection(collection)

	err := collection.FindOne(ctx, bson.D{{"email", r.Email}}).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			level.Error(logger).Log("error", "User not found")
			return "", erro.NewErrInvalidCredentials()
		}
		level.Error(logger).Log("error", err)
		return "", err
	}

	return resp.PwdHsh, nil
}
