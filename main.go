package main

import (
	"database/sql"
	"log"

	"github.com/babyfaceeasy/simplebank/api"
	db "github.com/babyfaceeasy/simplebank/db/sqlc"
	"github.com/babyfaceeasy/simplebank/util"
	_ "github.com/lib/pq"
)

func main()  {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Can not create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}
}