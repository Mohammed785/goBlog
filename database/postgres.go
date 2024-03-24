package database

import (
	"context"
	"log"

	"github.com/Mohammed785/goBlog/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToPostgres() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Config("DATABASE_URL"))
	if err != nil {
		log.Fatalln("couldn't connect to postgresql: ", err)
	}
	if err:=dbpool.Ping(context.Background());err!=nil{
		log.Fatalln(err)
	}
	return dbpool
}
