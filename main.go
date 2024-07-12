package main

import (
	"github/ariifysp/workshop-go/config"
	"github/ariifysp/workshop-go/databases"
	"github/ariifysp/workshop-go/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
