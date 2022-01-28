package user

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"

	erro "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/errors"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/utils"
	ent "github.com/fCalixto-Gb/Final-Project/HTTPService/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Get(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.GetUserResp{}, args.Error(1)
	}
	return result.(ent.GetUserResp), args.Error(1)
}

func (m *MockRepo) Create(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	return result.(ent.CreateUserResp), args.Error(1)
}

func (m *MockRepo) Delete(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.DeleteUserResp{}, args.Error(1)
	}
	return result.(ent.DeleteUserResp), args.Error(1)
}

func (m *MockRepo) Update(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.UpdateUserResp{}, args.Error(1)
	}
	return result.(ent.UpdateUserResp), args.Error(1)
}

func (m *MockRepo) Authenticate(ctx context.Context, r ent.AuthenticateReq) (ent.AuthenticateResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.AuthenticateResp{}, args.Error(1)
	}
	return result.(ent.AuthenticateResp), args.Error(1)
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
				Name:    "Francisco",
				Age:     20,
				Email:   "francisco.calixto@globant.com",
				Pwd:     "12345678",
				Country: "brazilian",
				Job:     "programmer",
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
				Name:    "Franco",
				Age:     32,
				Email:   "francisco.calixto@globant.com",
				Pwd:     "12345678",
				Country: "uruguayan",
				Job:     "cooker",
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
			testName: "uinvalid arguments",
			request: ent.CreateUserReq{
				Name:    "Franco",
				Age:     32,
				Email:   "francisco.calixto@globant.com",
				Pwd:     "12345678",
				Country: "uruguayan",
				Job:     "",
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
			mockRepo.On("Create", ctx, tc.request).Return(repoRes, err)
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
		Name:    "Francisco",
		Age:     20,
		Email:   "francisco.calixto@globant.com",
		Job:     "programmer",
		Country: "brazilian",
		Created: date,
		Id:      validID,
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
			mockRepo.On("Get", ctx, tc.request).Return(repoRes, err)
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
			mockRepo.On("Delete", ctx, tc.request).Return(repoRes, err)
			res, err := srvc.DeleteUser(ctx, tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}
