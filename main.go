package main

import (
	"fmt"
	"sushi-api/config"
	"sushi-api/databases"
	"sushi-api/server"
)

func main() {
	conf := config.GetConfig()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	fmt.Println("Starting server on port", conf.Server.Port)
	server.Start()
}
