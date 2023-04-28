package http1

import (
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type CreateUser struct {
	Name        string `json:"name,omitempty" db:"name, omitempty"`
	Description string `json:"description,omitempty" db:"description, omitempty"`
	MetaData    string `json:"metadata,omitempty" db:"metadata, omitempty"`
}
type UpdateUser struct {
	Name        string `json:"name,omitempty" db:"name, omitempty"`
	Description string `json:"description,omitempty" db:"description, omitempty"`
	MetaData    string `json:"metadata,omitempty" db:"metadata, omitempty"`
}

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
		encodeCreateRequest,
	))
	r.Handle("/update", httptransport.NewServer(
		UpdateHandler(db),
		decodeUpdateRequest,
		encodeUpdateRequest,
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
	var req CreateUser

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	fmt.Println(req.Name, req.MetaData, req.Description)
	return req, nil
}
func encodeCreateRequest(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUser

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	fmt.Println(req.Name, req.MetaData, req.Description)
	return req, nil
}
func encodeUpdateRequest(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
