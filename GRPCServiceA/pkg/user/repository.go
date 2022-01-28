package user

import (
	"context"
	"database/sql"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
	erro "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/errors"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/utils"
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(logger log.Logger, db *sql.DB) *repo {
	return &repo{
		db:     db,
		logger: logger,
	}
}

// Get user from database by id or email
func (repo *repo) GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error) {
	logger := log.With(repo.logger, "Repository method", "GetUser")
	var resp ent.GetUserResp
	getQuery := utils.GetQuery(r.Id) // gets corresponding query for Id or Email

	p, err := repo.db.PrepareContext(ctx, getQuery)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.GetUserResp{}, err
	}
	defer p.Close()

	err = p.QueryRowContext(ctx, r.Id).Scan(&resp.Id, &resp.Name, &resp.Age, &resp.Email, &resp.Country, &resp.Job, &resp.Created)
	if err != nil {
		if err == sql.ErrNoRows { // no rows matched -> no user found
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

	// check (by email) if user already exists in the database
	if _, err := repo.GetUser(ctx, ent.GetUserReq{Id: user.Email}); err == nil {
		level.Error(logger).Log("error", "User already exists")
		return ent.CreateUserResp{}, erro.NewErrAlreadyExists()
	}

	p, err := repo.db.PrepareContext(ctx, utils.CreateQuery)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.CreateUserResp{}, err
	}
	defer p.Close()

	_, err = p.ExecContext(ctx, user.Id, user.Name, user.Age, user.Email, user.PwdHsh, user.Country, user.Job, user.Created)
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
	deleteQuery := utils.DeleteQuery(r.Id) // gets corresponding query for Id or Email

	p, err := repo.db.PrepareContext(ctx, deleteQuery)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.DeleteUserResp{}, err
	}
	defer p.Close()

	res, err := p.ExecContext(ctx, r.Id)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.DeleteUserResp{}, err
	}
	if rows, _ := res.RowsAffected(); rows != 1 {
		level.Error(logger).Log("error", "User not found")
		return ent.DeleteUserResp{}, erro.NewErrUserNotFound()
	}
	return ent.DeleteUserResp{
		Deleted: utils.TimeNow(),
	}, nil
}

// Create new user in the database
func (repo *repo) UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error) {
	logger := log.With(repo.logger, "Repository method", "UpdateUser")

	// check (by email) if user exists in the database
	_, err := repo.GetUser(ctx, ent.GetUserReq{Id: r.Email})
	if _, ok := err.(*erro.ErrUserNotFound); ok {
		return ent.UpdateUserResp{}, err
	}

	p, err := repo.db.PrepareContext(ctx, utils.UpdateQuery)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.UpdateUserResp{}, err
	}
	defer p.Close()

	_, err = p.ExecContext(ctx, r.Name, r.Age, r.Email, r.Pwd, r.Country, r.Job, r.Email)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.UpdateUserResp{}, err
	}
	return ent.UpdateUserResp{
		Updated: utils.TimeNow(),
	}, nil
}

func (repo *repo) AuthenticateUser(ctx context.Context, r ent.AuthenticateReq) (string, error) {
	logger := log.With(repo.logger, "Repository method", "AuthenticateUser")
	var pwdhsh string

	p, err := repo.db.PrepareContext(ctx, utils.AuthenticateQuery)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return "", err
	}
	defer p.Close()

	err = p.QueryRowContext(ctx, r.Email).Scan(&pwdhsh)
	if err != nil {
		if err == sql.ErrNoRows { // no rows matched -> no user found
			level.Error(logger).Log("error", "User not found")
			return "", erro.NewErrInvalidCredentials()
		}
		level.Error(logger).Log("error", err)
		return "", err
	}
	return pwdhsh, nil
}
