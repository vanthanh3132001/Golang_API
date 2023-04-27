package main

import (
	"fmt"
	"github.com/example/dashBoard/api/http1"
	"github.com/example/dashBoard/postgresDB"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
)

func main() {
	// database connect
	var cfg postgresDB.Config
	loadConfig(&cfg)
	var sql = new(postgresDB.SQL)
	sql.Connect(cfg)
	defer sql.Close()

	var handler http.Handler
	handler = http1.MakeHandler(sql.Db)

	// Khởi động HTTP server với đối tượng xử lý request đã tạo
	startHTTPServer(handler)

}
func startHTTPServer(handler http.Handler) {
	// Khởi động HTTP server trên cổng 8080
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
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
