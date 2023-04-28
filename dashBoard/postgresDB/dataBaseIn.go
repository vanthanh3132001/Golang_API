package postgresDB

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	SaveUser(context context.Context, user MyRequest, db *sqlx.DB) (MyRequest, error)
}
