package user

import (
	"context"
	"log"
	"net"
	"testing"

	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
	erro "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/errors"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/user/proto"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// This file tests transport layer + endpoints by mocking the service layer, with the help of google.golang.org/grpc/test/bufconn
// in order to avoid opening any ports but still testing the layer functionality

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func BufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

type MockService struct {
	mock.Mock
}

func (m *MockService) GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.GetUserResp{}, args.Error(1)
	}
	return result.(ent.GetUserResp), args.Error(1)
}

func (m *MockService) CreateUser(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	return result.(ent.CreateUserResp), args.Error(1)
}

func (m *MockService) DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.DeleteUserResp{}, args.Error(1)
	}
	return result.(ent.DeleteUserResp), args.Error(1)
}

func (m *MockService) UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error) {
	args := m.Called(ctx, r)
	result := args.Get(0)
	if result == nil {
		return ent.UpdateUserResp{}, args.Error(1)
	}
	return result.(ent.UpdateUserResp), args.Error(1)
}

var validUser = ent.GetUserResp{
	Id:          validID,
	Name:        "Francisco",
	Age:         20,
	Nationality: "brazilian",
	Job:         "programmer",
	Email:       "francisco.calixto@globant.com",
	Created:     utils.TimeNow(),
}

func TestGetUserSuccess(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("GetUser", mock.Anything, ent.GetUserReq{Id: validID}).Return(validUser, nil)
	client := proto.NewUserServicesClient(conn)
	user, _ := client.GetUser(ctx, &proto.GetUserReq{Id: validID})

	// Test response
	assert.Equal(t, user.Error.Code, int32(0))
	assert.Equal(t, user.Error.Message, "ok")
}

func TestGetUserFail(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("GetUser", mock.Anything, ent.GetUserReq{Id: invalidID}).Return(ent.GetUserResp{}, erro.NewErrUserNotFound())
	client := proto.NewUserServicesClient(conn)
	user, _ := client.GetUser(ctx, &proto.GetUserReq{Id: invalidID})

	// Test response
	assert.Equal(t, user.Error.Code, int32(5))
	assert.Equal(t, user.Error.Message, "user not found")
}

func TestCreateUserSuccess(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("CreateUser", mock.Anything, ent.CreateUserReq{
		Name:        "Francisco",
		Age:         20,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "12345678",
		Nationality: "brazilian",
		Job:         "programmer",
	}).Return(ent.CreateUserResp{
		Id:      validID,
		Created: utils.TimeNow()}, nil)

	client := proto.NewUserServicesClient(conn)
	user, _ := client.CreateUser(ctx, &proto.CreateUserReq{
		Name:        "Francisco",
		Age:         20,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "12345678",
		Nationality: "brazilian",
		Job:         "programmer",
	})

	// Test response
	assert.Equal(t, user.Error.Code, int32(0))
	assert.Equal(t, user.Error.Message, "ok")
}

func TestCreateUserFail(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("CreateUser", mock.Anything, ent.CreateUserReq{
		Name:        "Franco",
		Age:         20,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "87654321",
		Nationality: "uruguayan",
		Job:         "cooker",
	}).Return(ent.CreateUserResp{}, erro.NewErrAlreadyExists())

	client := proto.NewUserServicesClient(conn)
	user, _ := client.CreateUser(ctx, &proto.CreateUserReq{
		Name:        "Franco",
		Age:         20,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "87654321",
		Nationality: "uruguayan",
		Job:         "cooker",
	})

	// Test response
	assert.Equal(t, user.Error.Code, int32(6))
	assert.Equal(t, user.Error.Message, "user already exists")
}

func TestCreateUserInvalidArguments(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("CreateUser", mock.Anything, ent.CreateUserReq{
		Name:        "Franco",
		Age:         20,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "87654321",
		Nationality: "",
		Job:         "cooker",
	}).Return(ent.CreateUserResp{}, erro.NewErrInvalidArguments())

	client := proto.NewUserServicesClient(conn)
	user, _ := client.CreateUser(ctx, &proto.CreateUserReq{
		Name:        "Franco",
		Age:         20,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "87654321",
		Nationality: "",
		Job:         "cooker",
	})

	// Test response
	assert.Equal(t, user.Error.Code, int32(3))
	assert.Equal(t, user.Error.Message, "invalid argument(s)")
}

func TestDeleteUserSuccess(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("DeleteUser", mock.Anything, ent.DeleteUserReq{Id: validID}).Return(ent.DeleteUserResp{Deleted: utils.TimeNow()}, nil)

	client := proto.NewUserServicesClient(conn)
	user, _ := client.DeleteUser(ctx, &proto.DeleteUserReq{
		Id: validID,
	})

	// Test response
	assert.Equal(t, user.Error.Code, int32(0))
	assert.Equal(t, user.Error.Message, "ok")
}

func TestDeleteUserFail(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("DeleteUser", mock.Anything, ent.DeleteUserReq{Id: invalidID}).Return(ent.DeleteUserResp{}, erro.NewErrUserNotFound())

	client := proto.NewUserServicesClient(conn)
	user, _ := client.DeleteUser(ctx, &proto.DeleteUserReq{
		Id: invalidID,
	})

	// Test response
	assert.Equal(t, user.Error.Code, int32(5))
	assert.Equal(t, user.Error.Message, "user not found")
}

func TestUpdateUserSuccess(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	validUpdateReq := ent.UpdateUserReq{
		Id:          validID,
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         utils.HashPwd("12345678"),
		Nationality: "brazilian",
		Job:         "programmer",
	}

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("UpdateUser", mock.Anything, validUpdateReq).Return(ent.UpdateUserResp{Updated: utils.TimeNow()}, nil)

	client := proto.NewUserServicesClient(conn)
	user, _ := client.UpdateUser(ctx, &proto.UpdateUserReq{
		Id:          validID,
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         utils.HashPwd("12345678"),
		Nationality: "brazilian",
		Job:         "programmer",
	})

	// Test response
	assert.Equal(t, user.Error.Code, int32(0))
	assert.Equal(t, user.Error.Message, "ok")
}

func TestUpdateUserInvalidArgs(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	invalidUpdateReq := ent.UpdateUserReq{
		Id:          validID,
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "",
		Nationality: "brazilian",
		Job:         "programmer",
	}

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("UpdateUser", mock.Anything, invalidUpdateReq).Return(ent.UpdateUserResp{}, erro.NewErrInvalidArguments())

	client := proto.NewUserServicesClient(conn)
	user, _ := client.UpdateUser(ctx, &proto.UpdateUserReq{
		Id:          validID,
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "",
		Nationality: "brazilian",
		Job:         "programmer",
	})

	// Test response
	assert.Equal(t, user.Error.Code, int32(3))
	assert.Equal(t, user.Error.Message, "invalid argument(s)")
}

func TestUpdateUserNotFound(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	// create a gRPC server object (like in main.go)
	mockSvc := new(MockService)
	epts := MakeEndpoints(mockSvc)
	grpcServer := NewGRPCServer(epts)
	baseServer := grpc.NewServer()

	// Register before starting the service.
	proto.RegisterUserServicesServer(baseServer, grpcServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	invalidUpdateReq := ent.UpdateUserReq{
		Id:          invalidID,
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "",
		Nationality: "brazilian",
		Job:         "programmer",
	}

	// Setup the client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufDialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	mockSvc.AssertExpectations(t)
	mockSvc.On("UpdateUser", mock.Anything, invalidUpdateReq).Return(ent.UpdateUserResp{}, erro.NewErrUserNotFound())

	client := proto.NewUserServicesClient(conn)
	user, _ := client.UpdateUser(ctx, &proto.UpdateUserReq{
		Id:          invalidID,
		Name:        "Francisco",
		Age:         21,
		Email:       "francisco.calixto@globant.com",
		Pwd:         "",
		Nationality: "brazilian",
		Job:         "programmer",
	})

	// Test response
	assert.Equal(t, user.Error.Code, int32(5))
	assert.Equal(t, user.Error.Message, "user not found")
}
