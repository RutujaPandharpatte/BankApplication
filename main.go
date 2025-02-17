package main

import (
	"database/sql"
	"log"

	"github.com/RutujaPandharpatte/BankApplication/api"
	db "github.com/RutujaPandharpatte/BankApplication/db/sqlc"
	"github.com/RutujaPandharpatte/BankApplication/util"
	_ "github.com/lib/pq"
)

func main() {
	println("Starting main")

	config, err := util.LoadConfig(".") //config file is in current directory
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	println("Config loaded")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	println("DB connected")

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	println("Server created")

	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	println("Server started")
}
