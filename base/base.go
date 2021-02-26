package base

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-pg/pg/v10"
)

func Init() error {
	opt, err := pg.ParseURL("postgres://postgres:root@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	postgresDB = pg.Connect(opt)

	idNode, err = snowflake.NewNode(1)
	if err != nil {
		return err
	}

	return nil
	//TODO: remember to close connection while stopping service
}

func GetIDGenerator() *snowflake.Node {
	return idNode
}

func GetPostgresDB() *pg.DB {
	return postgresDB
}
