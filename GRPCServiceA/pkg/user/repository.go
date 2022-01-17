package user

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"

	ent "github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/entities"
	erro "github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/errors"
	"github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/utils"
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

	err = p.QueryRowContext(ctx, r.Id).Scan(&resp.Id, &resp.Name, &resp.Age, &resp.Email, &resp.Nationality, &resp.Job, &resp.Created)
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
func (repo *repo) CreateUser(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error) {
	logger := log.With(repo.logger, "Repository method", "GetUser")

	// check (by email) if user already exists in the database
	if _, err := repo.GetUser(ctx, ent.GetUserReq{Id: r.Email}); err == nil {
		level.Error(logger).Log("error", "User already exists")
		return ent.CreateUserResp{}, erro.NewErrAlreadyExists()
	}

	id := utils.NewId()
	created := utils.TimeNow()
	pwdHsh := utils.HashPwd(r.Pwd)

	p, err := repo.db.PrepareContext(ctx, utils.CreateQuery)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.CreateUserResp{}, err
	}
	defer p.Close()

	_, err = p.ExecContext(ctx, id, r.Name, r.Age, r.Email, pwdHsh, r.Nationality, r.Job, created)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return ent.CreateUserResp{}, err
	}
	return ent.CreateUserResp{
		Id:      id,
		Created: created,
	}, nil
}
