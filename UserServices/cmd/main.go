package main

import (
	"context"

	"net"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/user"
	"github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/user/proto"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/runtime/protoiface"
)

var (
	grpcServerEndpoint = "localhost:50000"
	httpListen         = "localhost:8080"
	// dbCredentials      = "root:root"
	dbAddr   = "localhost:27017" // mongo is the container name
	database = "Users"
	applyURI = "mongodb://" + dbAddr
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "service", "gRPCServiceA")
	ctx := context.Background()

	// start db connection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(applyURI))
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		panic(err.Error())
	}
	// Ping to check if connection is completed
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		level.Info(logger).Log("error", "Unable to ping Mongodb")
		os.Exit(1)
	}
	db := client.Database(database)

	// initialize layers
	repo := user.NewRepo(logger, db)
	svc := user.NewService(logger, repo)
	epts := user.MakeEndpoints(svc)
	grpcServer := user.NewGRPCServer(epts)

	grpcListener, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	// grpc server start
	baseServer := grpc.NewServer()
	proto.RegisterUserServicesServer(baseServer, grpcServer)
	level.Info(logger).Log("msg", "Server started")
	go baseServer.Serve(grpcListener)

	// mux listener - acts as grpc client
	mux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(httpResponseModifier),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = proto.RegisterUserServicesHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	http.Handle("/", mux)
	level.Info(logger).Log("msg", "started listening 8080")
	http.ListenAndServe(httpListen, nil)
}

func httpResponseModifier(ctx context.Context, w http.ResponseWriter, _ protoiface.MessageV1) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code
	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		w.WriteHeader(code)
	}

	return nil
}
