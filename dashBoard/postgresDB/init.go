package postgresDB

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Config struct {
	Server struct {
		Port       string `yaml:"port"`
		Host       string `yaml:"host"`
		JwtExpires int    `yaml:"jwtExpires"`
	}
	Database struct {
		DbName     string `yaml:"dbName"`
		DbHost     string `yaml:"dbHost"`
		DbPort     string `yaml:"dbPort"`
		DbUserName string `yaml:"dbUserName"`
		DbPassword string `yaml:"dbPassword"`
	}
}
type SQL struct {
	Db *sqlx.DB
}

//Connect to postgres database
func (s *SQL) Connect(cfg Config) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.DbHost,
		cfg.Database.DbPort,
		cfg.Database.DbUserName,
		cfg.Database.DbPassword,
		cfg.Database.DbName)
	//fmt.Println(cfg.Database.DbHost,
	//	cfg.Database.DbPort,
	//	cfg.Database.DbUserName,
	//	cfg.Database.DbPassword,
	//	cfg.Database.DbName)
	s.Db = sqlx.MustConnect("postgres", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucess")

}

func (s *SQL) Close() {
	s.Db.Close()
}
