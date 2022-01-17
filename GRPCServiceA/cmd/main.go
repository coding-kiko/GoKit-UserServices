package main

import (
	"database/sql"
	"net"
	"os"

	_ "gopkg.in/go-sql-driver/mysql.v1"

	"github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/user"
	"github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/user/proto"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "service", "gRPCServiceA")

	var db *sql.DB
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/Users")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	repo := user.NewRepo(logger, db)
	svc := user.NewService(logger, repo)
	epts := user.MakeEndpoints(svc)
	grpcServer := user.NewGRPCServer(epts, logger)

	grpcListener, err := net.Listen("tcp", ":50000")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	baseServer := grpc.NewServer()
	proto.RegisterUserServicesServer(baseServer, grpcServer)
	level.Info(logger).Log("msg", "Server started")
	baseServer.Serve(grpcListener)
}
