package user

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
	erro "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/errors"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/utils"
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

var mockNewUser = &ent.User{
	Name:        "Francisco",
	Age:         20,
	PwdHsh:      utils.HashPwd("12345678"),
	Nationality: "brazilian",
	Job:         "programmer",
	Email:       "francisco.calixto@globant.com",
	Id:          validID,
	Created:     utils.TimeNow(),
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
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Nationality, mockNewUser.Job, mockNewUser.Created)

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
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Nationality, mockNewUser.Job, mockNewUser.Created)

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
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"})

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
				sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Nationality, mockNewUser.Job, mockNewUser.Created)

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
				sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Nationality, mockNewUser.Job, mockNewUser.Created)

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
				sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"})

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
		Id:          "francisco.calixto@globant.com",
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         utils.HashPwd("12345678"),
		Nationality: "brazilian",
		Job:         "programmer",
	}

	invalidFieldReq := ent.UpdateUserReq{
		Id:          validID,
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         utils.HashPwd("12345678"),
		Nationality: "",
		Job:         "programmer",
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
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"}).
					AddRow(validID, req.Name, req.Age, req.Email, req.Nationality, req.Job, utils.TimeNow())

				mock.ExpectPrepare(utils.GetQueryByEmail)
				mock.ExpectQuery(utils.GetQueryByEmail).
					WithArgs(req.Id).
					WillReturnRows(rows)
				mock.ExpectPrepare(utils.UpdateByEmail)
				mock.ExpectExec(utils.UpdateByEmail).
					WithArgs(req.Name, req.Age, req.Email, req.Pwd, req.Nationality, req.Job, req.Id).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			checkResponse: func(t *testing.T, resp ent.UpdateUserResp, err error) {
				assert.Nil(t, err)
			},
		},
		// {
		// 	testName: "user not found",
		// 	request:  invalidIDReq,
		// 	execute: func(mock sqlmock.Sqlmock, req ent.UpdateUserReq) {
		// 		rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"}).
		// 			AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Nationality, mockNewUser.Job, mockNewUser.Created)

		// 		mock.ExpectPrepare(utils.GetQueryByEmail)
		// 		mock.ExpectQuery(utils.GetQueryByEmail).
		// 			WithArgs(req.Email).
		// 			WillReturnRows(rows)
		// 		mock.ExpectPrepare(utils.UpdateByEmail)
		// 		mock.ExpectExec(utils.UpdateByEmail).
		// 			WithArgs(req.Name, req.Age, req.Email, req.Pwd, req.Nationality, req.Job, req.Email).
		// 			WillReturnResult(sqlmock.NewResult(0, 0))
		// 	},
		// 	checkResponse: func(t *testing.T, resp ent.UpdateUserResp, err error) {
		// 		assert.Equal(t, err.Error(), "user not found")
		// 		_, ok := err.(*erro.ErrUserNotFound)
		// 		assert.Equal(t, ok, true)
		// 	},
		// },
		{
			testName: "invalid arguments",
			request:  invalidFieldReq,
			execute: func(mock sqlmock.Sqlmock, req ent.UpdateUserReq) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "email", "nationality", "job", "created"}).
					AddRow(mockNewUser.Id, mockNewUser.Name, mockNewUser.Age, mockNewUser.Email, mockNewUser.Nationality, mockNewUser.Job, mockNewUser.Created)

				mock.ExpectPrepare(utils.GetQueryByEmail)
				mock.ExpectQuery(utils.GetQueryByEmail).
					WithArgs(req.Id).
					WillReturnRows(rows)
				mock.ExpectPrepare(utils.UpdateByEmail)
				mock.ExpectExec(utils.UpdateByEmail).
					WithArgs(req.Name, req.Age, req.Email, req.Pwd, "", req.Job, invalidID).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			checkResponse: func(t *testing.T, resp ent.UpdateUserResp, err error) {
				assert.Equal(t, err.Error(), "invalid argument(s)")
				_, ok := err.(*erro.ErrInvalidArguments)
				assert.Equal(t, ok, true)
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
