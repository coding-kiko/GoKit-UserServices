package user

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	ent "github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/entities"
	erro "github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/errors"
	"github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/utils"
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
