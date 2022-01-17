package main

import (
	"net/http"
	"os"

	"github.com/coding-kiko/GoKit-Project-Bootcamp/HTTPService/pkg/user"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcDial  = "localhost:50000"
	httpServe = "127.0.0.1:8000"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "service", "HTTPService")

	// grpc client connection
	conn, err := grpc.Dial(grpcDial, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Unable to dial Grpc")
	}
	defer conn.Close()

	repo := user.NewGRPClient(logger, conn)
	svc := user.NewService(logger, repo)
	epts := user.MakeEndpoints(svc)
	muxHandler := user.NewMuxApi(epts, logger)
	srv := &http.Server{
		Handler: muxHandler,
		Addr:    httpServe,
	}
	level.Info(logger).Log("msg", "Http Server started and listening")
	srv.ListenAndServe()
}
