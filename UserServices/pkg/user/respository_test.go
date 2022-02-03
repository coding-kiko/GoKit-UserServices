package user

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	ent "github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/entities"
	erro "github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/errors"
	"github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/utils"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/stretchr/testify/assert"
)

// Mock mysql database
func NewMock(logger log.Logger) (*sql.DB, sqlmock.Sqlmock) {
	// Open a mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		level.Error(logger).Log("error opening mock database connection", err)
	}
	return db, mock
}

var validID = utils.NewId()
var invalidID = "hd2h97643gg2g9d7dhjfj"

var mockNewUser = ent.User{
	Name:    "Francisco",
	Age:     20,
	PwdHsh:  utils.HashPwd("12345678"),
	Country: "brazilian",
	Job:     "programmer",
	Email:   "francisco.calixto@globant.com",
	Id:      validID,
	Created: utils.TimeNow(),
}

func TestRepoGetUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "repo_test",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	db, mock := NewMock(logger)
	defer db.Close()

	repo := NewRepo(logger, db)

	testCases := []struct {
		testName      string
		request       ent.GetUserReq
		execute       func(mock sqlmock.Sqlmock, req ent.GetUserReq)
		checkResponse func(t *testing.T, resp ent.GetUserResp, err error)
	}{
		{
			testName: "get user by email successfull",
			request: ent.GetUserReq{
				Id: "francisco.calixto@globant.com",
			},
			execute: func(mock sqlmock.Sqlmock, req ent.GetUserReq) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "country", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Country, mockNewUser.Job, mockNewUser.Created)

				mock.ExpectPrepare(utils.GetQueryByEmail)
				mock.ExpectQuery(utils.GetQueryByEmail).
					WithArgs(req.Id).
					WillReturnRows(rows)
			},
			checkResponse: func(t *testing.T, resp ent.GetUserResp, err error) {
				assert.NoError(t, err)
				assert.Equal(t, resp.Created, mockNewUser.Created)
				assert.Equal(t, resp.Email, mockNewUser.Email)
			},
		},
		{
			testName: "get user by id successfull",
			request: ent.GetUserReq{
				Id: validID,
			},
			execute: func(mock sqlmock.Sqlmock, req ent.GetUserReq) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "Country", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Country, mockNewUser.Job, mockNewUser.Created)

				mock.ExpectPrepare(utils.GetQueryById)
				mock.ExpectQuery(utils.GetQueryById).
					WithArgs(req.Id).
					WillReturnRows(rows)
			},
			checkResponse: func(t *testing.T, resp ent.GetUserResp, err error) {
				assert.NoError(t, err)
				assert.Equal(t, resp.Created, mockNewUser.Created)
				assert.Equal(t, resp.Id, validID)
			},
		},
		{
			testName: "get user fail",
			request: ent.GetUserReq{
				Id: invalidID,
			},
			execute: func(mock sqlmock.Sqlmock, req ent.GetUserReq) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "country", "job", "created"})

				mock.ExpectPrepare(utils.GetQueryById)
				mock.ExpectQuery(utils.GetQueryById).
					WithArgs(req.Id).
					WillReturnRows(rows)
			},
			checkResponse: func(t *testing.T, resp ent.GetUserResp, err error) {
				assert.Equal(t, err.Error(), "user not found")
				_, ok := err.(*erro.ErrUserNotFound)
				assert.Equal(t, true, ok)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()

			tc.execute(mock, tc.request)

			res, err := repo.GetUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestRepoDeleteUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "repo_test",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	db, mock := NewMock(logger)
	defer db.Close()

	repo := NewRepo(logger, db)

	testCases := []struct {
		testName      string
		request       ent.DeleteUserReq
		execute       func(mock sqlmock.Sqlmock, req ent.DeleteUserReq)
		checkResponse func(t *testing.T, resp ent.DeleteUserResp, err error)
	}{
		{
			testName: "delete user by id successfull",
			request: ent.DeleteUserReq{
				Id: validID,
			},
			execute: func(mock sqlmock.Sqlmock, req ent.DeleteUserReq) {
				sqlmock.NewRows([]string{"id", "name", "age", "email", "country", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Country, mockNewUser.Job, mockNewUser.Created)

				mock.ExpectPrepare(utils.DeleteById)
				mock.ExpectExec(utils.DeleteById).
					WithArgs(req.Id).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			checkResponse: func(t *testing.T, resp ent.DeleteUserResp, err error) {
				assert.NoError(t, err)
			},
		},
		{
			testName: "delete user by email successfull",
			request: ent.DeleteUserReq{
				Id: "francisco.calixto@globant.com",
			},
			execute: func(mock sqlmock.Sqlmock, req ent.DeleteUserReq) {
				sqlmock.NewRows([]string{"id", "name", "age", "email", "country", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Country, mockNewUser.Job, mockNewUser.Created)

				mock.ExpectPrepare(utils.DeleteByEmail)
				mock.ExpectExec(utils.DeleteByEmail).
					WithArgs(req.Id).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			checkResponse: func(t *testing.T, resp ent.DeleteUserResp, err error) {
				assert.NoError(t, err)
			},
		},
		{
			testName: "delete user not found",
			request: ent.DeleteUserReq{
				Id: invalidID,
			},
			execute: func(mock sqlmock.Sqlmock, req ent.DeleteUserReq) {
				sqlmock.NewRows([]string{"id", "name", "age", "email", "country", "job", "created"})

				mock.ExpectPrepare(utils.DeleteById)
				mock.ExpectExec(utils.DeleteById).
					WithArgs(req.Id).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			checkResponse: func(t *testing.T, resp ent.DeleteUserResp, err error) {
				assert.Equal(t, err.Error(), "user not found")
				_, ok := err.(*erro.ErrUserNotFound)
				assert.Equal(t, true, ok)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()

			tc.execute(mock, tc.request)

			res, err := repo.DeleteUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestRepoUpdateUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "repo_test",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	db, mock := NewMock(logger)
	defer db.Close()

	repo := NewRepo(logger, db)

	validUpdateReq := ent.UpdateUserReq{
		Name:    "Francisco",
		Age:     21,
		Email:   "francisco.calixto@globant.com",
		Pwd:     utils.HashPwd("12345678"),
		Country: "brazilian",
		Job:     "programmer",
	}

	testCases := []struct {
		testName      string
		request       ent.UpdateUserReq
		execute       func(mock sqlmock.Sqlmock, req ent.UpdateUserReq)
		checkResponse func(t *testing.T, resp ent.UpdateUserResp, err error)
	}{
		{
			testName: "update user successfully",
			request:  validUpdateReq,
			execute: func(mock sqlmock.Sqlmock, req ent.UpdateUserReq) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "country", "job", "created"}).
					AddRow(validID, req.Name, req.Age, req.Email, req.Country, req.Job, utils.TimeNow())

				mock.ExpectPrepare(utils.GetQueryByEmail)
				mock.ExpectQuery(utils.GetQueryByEmail).
					WithArgs(req.Email).
					WillReturnRows(rows)
				mock.ExpectPrepare(utils.UpdateQuery)
				mock.ExpectExec(utils.UpdateQuery).
					WithArgs(req.Name, req.Age, req.Email, req.Pwd, req.Country, req.Job, req.Email).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			checkResponse: func(t *testing.T, resp ent.UpdateUserResp, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()

			tc.execute(mock, tc.request)

			res, err := repo.UpdateUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestRepoCreateUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "repo_test",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	db, mock := NewMock(logger)
	defer db.Close()

	repo := NewRepo(logger, db)

	testCases := []struct {
		testName      string
		request       ent.User
		execute       func(mock sqlmock.Sqlmock, req ent.User)
		checkResponse func(t *testing.T, resp ent.CreateUserResp, err error)
	}{
		{
			testName: "create user already exists",
			request:  mockNewUser,
			execute: func(mock sqlmock.Sqlmock, req ent.User) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "country", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Country, mockNewUser.Job, mockNewUser.Created)

				mock.ExpectPrepare(utils.GetQueryByEmail)
				mock.ExpectQuery(utils.GetQueryByEmail).
					WithArgs(req.Email).
					WillReturnRows(rows)
				mock.ExpectPrepare(utils.CreateQuery)
				mock.ExpectExec(utils.CreateQuery).
					WithArgs(req.Id, req.Name, req.Age, req.Email, req.PwdHsh, req.Country, req.Job, req.Created).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			checkResponse: func(t *testing.T, resp ent.CreateUserResp, err error) {
				_, ok := err.(*erro.ErrAlreadyExists)
				assert.True(t, ok)

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()

			tc.execute(mock, tc.request)

			res, err := repo.CreateUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestRepoAuthenticateUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "repo_test",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	invalidCredentials := ent.AuthenticateReq{
		Email: "doesntexist",
		Pwd:   "admin",
	}

	db, mock := NewMock(logger)
	defer db.Close()

	repo := NewRepo(logger, db)

	testCases := []struct {
		testName      string
		request       ent.AuthenticateReq
		execute       func(mock sqlmock.Sqlmock, req ent.AuthenticateReq)
		checkResponse func(t *testing.T, resp string, err error)
	}{
		{
			testName: "authenticate user not found",
			request:  invalidCredentials,
			execute: func(mock sqlmock.Sqlmock, req ent.AuthenticateReq) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "country", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Country, mockNewUser.Job, mockNewUser.Created)

				mock.ExpectPrepare(utils.AuthenticateQuery)
				mock.ExpectQuery(utils.AuthenticateQuery).
					WithArgs(req.Email).
					WillReturnRows(rows)
			},
			checkResponse: func(t *testing.T, resp string, err error) {
				assert.Empty(t, resp)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()

			tc.execute(mock, tc.request)

			res, err := repo.AuthenticateUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}
