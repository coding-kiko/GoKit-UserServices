package main

import (
	"context"
	"database/sql"
	"flag"
	"net"
	"net/http"
	"os"

	_ "gopkg.in/go-sql-driver/mysql.v1"

	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/user"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/user/proto"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// this is temporary while I dont implement env vars
var (
	dataSourceName     = "root:@(127.0.0.1:3306)/Users"
	grpcPort           = ":50000"
	grpcServerEndpoint = "localhost:50000"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "service", "gRPCServiceA")
	ctx := context.Background()

	var db *sql.DB
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	repo := user.NewRepo(logger, db)
	svc := user.NewService(logger, repo)
	epts := user.MakeEndpoints(svc)
	grpcServer := user.NewGRPCServer(epts)

	grpcListener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	flag.Parse()

	baseServer := grpc.NewServer()
	proto.RegisterUserServicesServer(baseServer, grpcServer)
	level.Info(logger).Log("msg", "Server started")
	go baseServer.Serve(grpcListener)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = proto.RegisterUserServicesHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	http.Handle("/", mux)
	http.ListenAndServe(":8000", nil)
}
