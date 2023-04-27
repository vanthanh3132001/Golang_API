package postgresDB

import (
	"context"
	"errors"
	n "github.com/example/dashBoard/api/http1"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type myService struct {
	Id          string
	Name        string
	Description string
	MetaData    string
	UpdatedTime string
	UpdatedBy   string
}
type UserRepoImpl struct {
	sql *SQL
}

func (s *myService) GetByID(ctx context.Context, id string) (string, error) {
	return "", nil
	// implementation
}

func SaveUser(context context.Context, user n.MyRequest, db *sqlx.DB) (n.MyRequest, error) {

	statement := `
		INSERT INTO dashboard(id,name,description,metadata,updated_time,updated_by)
		VALUES (:id,:name,:description,:metadata,:updated_time,:updated_by)
	`

	_, err := db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errors.New("Ng dung ton tai")
			}
		}
		return user, errors.New("dang ky that bai")
	}
	return user, nil
}
