package user

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"

	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
	erro "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/errors"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.GetUserResp{}, args.Error(1)
	}
	return result.(ent.GetUserResp), args.Error(1)
}

func (m *MockRepo) CreateUser(ctx context.Context, user ent.User) (ent.CreateUserResp, error) {
	args := m.Called(ctx, user)
	result := args.Get(0)
	return result.(ent.CreateUserResp), args.Error(1)
}

func (m *MockRepo) DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.DeleteUserResp{}, args.Error(1)
	}
	return result.(ent.DeleteUserResp), args.Error(1)
}

func (m *MockRepo) UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.UpdateUserResp{}, args.Error(1)
	}
	return result.(ent.UpdateUserResp), args.Error(1)
}

// test service CreateUser
func TestCreateUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserServiceTesting",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	testCases := []struct {
		testName      string
		request       ent.CreateUserReq
		response      func(req ent.CreateUserReq) (ent.CreateUserResp, error)
		checkResponse func(t *testing.T, resp ent.CreateUserResp, e error)
	}{
		{
			testName: "user created successfully",
			request: ent.CreateUserReq{
				Name:        "Francisco",
				Age:         20,
				Email:       "francisco.calixto@globant.com",
				Pwd:         "12345678",
				Nationality: "brazilian",
				Job:         "programmer",
			},
			response: func(req ent.CreateUserReq) (ent.CreateUserResp, error) {
				return ent.CreateUserResp{
					Id:      utils.NewId(),
					Created: utils.TimeNow(),
				}, nil
			},
			checkResponse: func(t *testing.T, resp ent.CreateUserResp, e error) {
				assert.Nil(t, e)
			},
		},
		{
			testName: "user already exists",
			request: ent.CreateUserReq{
				Name:        "Franco",
				Age:         32,
				Email:       "francisco.calixto@globant.com",
				Pwd:         "12345678",
				Nationality: "uruguayan",
				Job:         "cooker",
			},
			response: func(req ent.CreateUserReq) (ent.CreateUserResp, error) {
				return ent.CreateUserResp{}, erro.NewErrAlreadyExists()
			},
			checkResponse: func(t *testing.T, resp ent.CreateUserResp, e error) {
				assert.Empty(t, resp)
				_, ok := e.(*erro.ErrAlreadyExists)
				assert.EqualValues(t, ok, true)
				assert.Equal(t, e.Error(), "user already exists")
			},
		},
		{
			testName: "invalid arguments",
			request: ent.CreateUserReq{
				Name:        "Michael",
				Age:         24,
				Email:       "michael.scott@gmail.com",
				Pwd:         "admin",
				Nationality: "american",
				Job:         "",
			},
			response: func(req ent.CreateUserReq) (ent.CreateUserResp, error) {
				return ent.CreateUserResp{}, erro.NewErrInvalidArguments()
			},
			checkResponse: func(t *testing.T, resp ent.CreateUserResp, e error) {
				assert.Empty(t, resp)
				_, ok := e.(*erro.ErrInvalidArguments)
				assert.EqualValues(t, ok, true)
				assert.Equal(t, e.Error(), "invalid argument(s)")
			},
		},
	}

	mockRepo := new(MockRepo)
	mockRepo.AssertExpectations(t)
	srvc := NewService(logger, mockRepo)

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			repoRes, err := tc.response(tc.request)
			mockRepo.On("CreateUser", ctx, tc.request).Return(repoRes, err)
			res, err := srvc.CreateUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}

// Test service GetUser

func TestGetUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserServiceTesting",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	validID := utils.NewId()
	date := utils.TimeNow()
	validUser := ent.GetUserResp{
		Name:        "Francisco",
		Age:         20,
		Email:       "francisco.calixto@globant.com",
		Job:         "programmer",
		Nationality: "brazilian",
		Created:     date,
		Id:          validID,
	}

	testCases := []struct {
		testName      string
		request       ent.GetUserReq
		response      func(req ent.GetUserReq) (ent.GetUserResp, error)
		checkResponse func(t *testing.T, resp ent.GetUserResp, e error)
	}{
		{
			testName: "get user succesfull",
			request: ent.GetUserReq{
				Id: validID,
			},
			response: func(req ent.GetUserReq) (ent.GetUserResp, error) {
				return validUser, nil
			},
			checkResponse: func(t *testing.T, resp ent.GetUserResp, e error) {
				assert.Nil(t, e)
			},
		},
		{
			testName: "get inexistent user",
			request: ent.GetUserReq{
				Id: "not_a_valid_id",
			},
			response: func(req ent.GetUserReq) (ent.GetUserResp, error) {
				return ent.GetUserResp{}, erro.NewErrUserNotFound()
			},
			checkResponse: func(t *testing.T, resp ent.GetUserResp, e error) {
				_, ok := e.(*erro.ErrUserNotFound)
				assert.EqualValues(t, true, ok)
				assert.Empty(t, resp)
				assert.Equal(t, e.Error(), "user not found")
			},
		},
	}

	mockRepo := new(MockRepo)
	mockRepo.AssertExpectations(t)
	srvc := NewService(logger, mockRepo)

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			repoRes, err := tc.response(tc.request)
			mockRepo.On("GetUser", ctx, tc.request).Return(repoRes, err)
			res, err := srvc.GetUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}

// Test service DeleteUser

func TestDeleteUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserServiceTesting",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	validID := utils.NewId()

	testCases := []struct {
		testName      string
		request       ent.DeleteUserReq
		response      func(req ent.DeleteUserReq) (ent.DeleteUserResp, error)
		checkResponse func(t *testing.T, resp ent.DeleteUserResp, e error)
	}{
		{
			testName: "delete user succesfull",
			request: ent.DeleteUserReq{
				Id: validID,
			},
			response: func(req ent.DeleteUserReq) (ent.DeleteUserResp, error) {
				return ent.DeleteUserResp{Deleted: utils.TimeNow()}, nil
			},
			checkResponse: func(t *testing.T, resp ent.DeleteUserResp, e error) {
				assert.Nil(t, e)
			},
		},
		{
			testName: "delete inexistent user",
			request: ent.DeleteUserReq{
				Id: "not_a_valid_id",
			},
			response: func(req ent.DeleteUserReq) (ent.DeleteUserResp, error) {
				return ent.DeleteUserResp{}, erro.NewErrUserNotFound()
			},
			checkResponse: func(t *testing.T, resp ent.DeleteUserResp, e error) {
				_, ok := e.(*erro.ErrUserNotFound)
				assert.EqualValues(t, true, ok)
				assert.Empty(t, resp)
				assert.Equal(t, e.Error(), "user not found")
			},
		},
	}

	mockRepo := new(MockRepo)
	mockRepo.AssertExpectations(t)
	srvc := NewService(logger, mockRepo)

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			repoRes, err := tc.response(tc.request)
			mockRepo.On("DeleteUser", ctx, tc.request).Return(repoRes, err)
			res, err := srvc.DeleteUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserServiceTesting",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	validUpdateReq := ent.UpdateUserReq{
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "12345678",
		Nationality: "brazilian",
		Job:         "programmer",
	}
	invalidUpdateReq := ent.UpdateUserReq{
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "12345678",
		Nationality: "brazilian",
		Job:         "programmer",
	}
	invalidArgsUpdateReq := ent.UpdateUserReq{
		Name:        "Francisco",
		Age:         21,
		Email:       "",
		Pwd:         "12345678",
		Nationality: "brazilian",
		Job:         "programmer",
	}

	testCases := []struct {
		testName      string
		request       ent.UpdateUserReq
		response      func(req ent.UpdateUserReq) (ent.UpdateUserResp, error)
		checkResponse func(t *testing.T, resp ent.UpdateUserResp, e error)
	}{
		{
			testName: "update user succesfull",
			request:  validUpdateReq,
			response: func(req ent.UpdateUserReq) (ent.UpdateUserResp, error) {
				return ent.UpdateUserResp{Updated: utils.TimeNow()}, nil
			},
			checkResponse: func(t *testing.T, resp ent.UpdateUserResp, e error) {
				assert.Nil(t, e)
			},
		},
		{
			testName: "update user not found",
			request:  invalidUpdateReq,
			response: func(req ent.UpdateUserReq) (ent.UpdateUserResp, error) {
				return ent.UpdateUserResp{}, erro.NewErrUserNotFound()
			},
			checkResponse: func(t *testing.T, resp ent.UpdateUserResp, e error) {
				_, ok := e.(*erro.ErrUserNotFound)
				assert.Equal(t, ok, true)
				assert.Empty(t, resp)
			},
		},
		{
			testName: "update user invalid arguments",
			request:  invalidArgsUpdateReq,
			response: func(req ent.UpdateUserReq) (ent.UpdateUserResp, error) {
				return ent.UpdateUserResp{}, erro.NewErrInvalidArguments()
			},
			checkResponse: func(t *testing.T, resp ent.UpdateUserResp, e error) {
				_, ok := e.(*erro.ErrInvalidArguments)
				assert.Equal(t, ok, true)
				assert.Empty(t, resp)
			},
		},
	}

	mockRepo := new(MockRepo)
	mockRepo.AssertExpectations(t)
	srvc := NewService(logger, mockRepo)

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			repoRes, err := tc.response(tc.request)
			mockRepo.On("UpdateUser", ctx, tc.request).Return(repoRes, err)
			res, err := srvc.UpdateUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}
