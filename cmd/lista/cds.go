package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CristianMarsico/seminario_Go/internal/config"
	"github.com/CristianMarsico/seminario_Go/internal/database"
	"github.com/CristianMarsico/seminario_Go/internal/service/lista"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	conf := config.LoadConfig(*configFile)

	db, err := database.NewDatabase(conf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	service, _ := lista.New(db, conf)

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	httpService := lista.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)

	r.Run()
}

func createSchema(db *sqlx.DB) error {

	schema1 := (`CREATE TABLE IF NOT EXISTS lista (
		id integer primary key autoincrement,
		name varchar(56) NOT NULL UNIQUE);`)

	_, err := db.Exec(schema1)
	if err != nil {
		return err
	}
	return nil
}
