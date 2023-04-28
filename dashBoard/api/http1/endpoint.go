package http1

import (
	"context"
	"fmt"
	ureq "github.com/example/dashBoard/postgresDB"
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
		req := request.(CreateUser)
		fmt.Println(req.Name, req.MetaData, req.Description)
		userId, _ := uuid.NewUUID()
		//now := time.Now()
		user := ureq.MyRequest{

			Id:          userId.String(),
			Name:        req.Name,
			Description: req.Description,
			MetaData:    req.MetaData,
			UpdatedTime: req.MetaData,
			UpdatedBy:   userId.String(),
		}
		user, _ = ureq.SaveUser(context, user, db)

		return HelloWorldResponse{Message: "TC create"}, nil
	}
}

func UpdateHandler(db *sqlx.DB) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUser)
		fmt.Println("req.Name, req.MetaData, req.Description")
		fmt.Println(req.Name, req.MetaData, req.Description)

		//now := time.Now()
		user2 := ureq.MyRequest{
			Id:          "1",
			Name:        req.Name,
			Description: req.Description,
			MetaData:    req.MetaData,
			UpdatedTime: "1",
			UpdatedBy:   "2",
		}
		user2, _ = ureq.UpdateUser(context, user2, db)

		return HelloWorldResponse{Message: "TC update"}, nil
	}
}
