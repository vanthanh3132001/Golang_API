package http

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

func MyEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// ép kiểu request từ interface{} sang struct MyRequest
		req := request.(MyRequest)

		if req.Name == "" {
			return nil, errors.New("name is required")
		}

		// trả về message và error là nil
		return MyResponse{Message: "Hello " + req.Name + "!"}, nil
	}
}
