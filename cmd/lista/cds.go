package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CristianMarsico/seminario_Go/internal/config"
	"github.com/CristianMarsico/seminario_Go/internal/database"
	"github.com/CristianMarsico/seminario_Go/internal/service/lista"
)

func main() {

	cfg := readConfig()

	db, err := database.NewDataBase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := lista.New(db, cfg)
	for _, m := range service.FindAll() {
		fmt.Println(m)
	}

}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
	}
	return cfg
}
