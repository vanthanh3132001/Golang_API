package postgresDB

import (
	"context"
	n "github.com/example/dashBoard/api/http1"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	SaveUser(context context.Context, user n.MyRequest, db *sqlx.DB) (n.MyRequest, error)
}
