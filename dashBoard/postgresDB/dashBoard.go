package postgresDB

import (
	"context"
	"errors"
	"fmt"
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

type MyRequest struct {
	Id          string `json:"id,omitempty" db:"id, omitempty"`
	Name        string `json:"name,omitempty" db:"name, omitempty"`
	Description string `json:"description,omitempty" db:"description, omitempty"`
	MetaData    string `json:"metadata,omitempty" db:"metadata, omitempty"`
	UpdatedTime string `json:"updated_time,omitempty" db:"updated_time, omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty" db:"updated_by, omitempty"`
}
type UpdateRequest struct {
	Name        string `json:"name,omitempty" db:"name, omitempty"`
	Description string `json:"description,omitempty" db:"description, omitempty"`
	MetaData    string `json:"metadata,omitempty" db:"metadata, omitempty"`
}

func SaveUser(context context.Context, user MyRequest, db *sqlx.DB) (MyRequest, error) {

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

func UpdateUser(context context.Context, user MyRequest, db *sqlx.DB) (MyRequest, error) {

	statement := `
		UPDATE dashboard SET  name = :name,description = :description  WHERE metadata = :metadata;
	`
	fmt.Println("****")
	fmt.Println(user.Name)
	fmt.Println(user.MetaData)
	fmt.Println(user.Description)
	fmt.Println("****")

	_, err := db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errors.New("Ng dung ton tai")
			}
		}
		return user, errors.New("cap nhat that bai")
	}
	return user, nil
}
