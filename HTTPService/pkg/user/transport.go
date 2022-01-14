package user

import (
	"context"
	"encoding/json"
	"net/http"

	erro "github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/errors"
	ent "github.com/coding-kiko/GoKit-Project-Bootcamp/HTTPService/pkg/entities"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func NewMuxApi(endpoints Endpoints, logger log.Logger) http.Handler {
	// options to support error control
	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorLogger(logger),
		kitHttp.ServerErrorEncoder(encodeErrorResponse), // encode error to the http.ResponseWriter whenever one is found processing a request
	}
	r := mux.NewRouter()
	// GetUser method handler
	r.Methods("GET").Path("/api/users/{Id}").Handler(
		kitHttp.NewServer(
			endpoints.GetUser,
			decodeGetUser,
			encodeResp,
			options...,
		),
	)
	// CreateUser method handler
	r.Methods("POST").Path("/api/users").Handler(
		kitHttp.NewServer(
			endpoints.CreateUser,
			decodeCreateUser,
			encodeResp,
			options...,
		),
	)
	return r
}

// Encodes to json to give a http response
func encodeResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

// decode request entering the http server
func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ent.GetUserReq
	// Getting passed id from path params
	Id := mux.Vars(r)["Id"]
	req.Id = Id
	return req, nil
}

// decode request entering the http server
func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ent.CreateUserReq
	// translating json to my CreateUserReq struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// encode my custom errors to json response
func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(erro.ErrToHTTPStatus(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
