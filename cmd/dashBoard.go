package main

import (
	"fmt"
	"github.com/example/dashBoard/postgresDB"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {
	// database connect
	var cfg postgresDB.Config
	loadConfig(&cfg)
	var sql = new(postgresDB.SQL)
	sql.Connect(cfg)
	defer sql.Close()
}

func loadConfig(cfg *postgresDB.Config) {
	f, err := os.Open("../env.dev.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
}
