package http1

import (
	"context"
	ureq "github.com/example/dashBoard/postgresDB"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	UserRepo ureq.UserRepo
}

func HelloWorldHandler() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloWorldRequest)
		message := "Hello, " + req.Name + "!!!"

		return HelloWorldResponse{Message: message}, nil
	}
}

func CreateHandler(db *sqlx.DB) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := ureq.CreateUser{}
		userId, _ := uuid.NewUUID()
		now := time.Now()
		user := MyRequest{

			Id:          userId.String(),
			Name:        req.Name,
			Description: req.Description,
			MetaData:    req.MetaData,
			UpdatedTime: now,
			UpdatedBy:   userId.String(),
		}
		user, _ = ureq.SaveUser(context, user, db)

		return HelloWorldResponse{Message: "TC"}, nil
	}
}
