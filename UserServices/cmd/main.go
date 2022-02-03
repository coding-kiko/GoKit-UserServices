package main

import (
	"context"
	"database/sql"
	"flag"
	"net"
	"net/http"
	"os"
	"strconv"

	_ "gopkg.in/go-sql-driver/mysql.v1"

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
	grpcServerEndpoint = flag.String("grpcServerEndpoint", "localhost:50000", "endpoint for grpc connection client <-> server")
	httpListen         = flag.String("httpAddr", "localhost:8000", "http listener address")
	dbCredentials      = flag.String("dbCredentials", "root:", "username and password to generate mysql connection")
	dbAddr             = flag.String("dbAddr", "localhost:3306", "mysql address")
	database           = flag.String("database", "Users", "database name")
	dataSourceName     = *dbCredentials + "@(" + *dbAddr + ")/" + *database
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "service", "gRPCServiceA")
	ctx := context.Background()

	flag.Parse()

	// start db connection
	var db *sql.DB
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// initialize layers
	repo := user.NewRepo(logger, db)
	svc := user.NewService(logger, repo)
	epts := user.MakeEndpoints(svc)
	grpcServer := user.NewGRPCServer(epts)

	grpcListener, err := net.Listen("tcp", *grpcServerEndpoint)
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
	err = proto.RegisterUserServicesHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	http.Handle("/", mux)
	http.ListenAndServe(*httpListen, nil)
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
