package http1

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func MakeHandler(db *sqlx.DB) http.Handler {
	r := http.NewServeMux()

	r.Handle("/hello", httptransport.NewServer(
		HelloWorldHandler(),
		decodeHelloWorldRequest,
		encodeHelloWorldResponse,
	))
	r.Handle("/create", httptransport.NewServer(
		CreateHandler(db),
		decodeCreateRequest,
		iecodeCreateRequest,
	))

	// Khởi động server trên cổng 8080
	http.ListenAndServe(":8080", r)
	return r
}

func decodeHelloWorldRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req HelloWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
func encodeHelloWorldResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req MyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
func iecodeCreateRequest(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
