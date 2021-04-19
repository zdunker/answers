package base

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
)

func Init() error {
	opt, err := pg.ParseURL("postgres://postgres:root@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	postgresDB = pg.Connect(opt)

	redisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	validate = validator.New()

	if idNode, err = snowflake.NewNode(1); err != nil {
		return err
	}

	if serverLocation, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		return err
	}

	sessionStore = sessions.NewFilesystemStore("../..", []byte("random"))

	return nil
	//TODO: remember to close connection while stopping service
}

func GetIDGenerator() *snowflake.Node {
	return idNode
}

func GetPostgresDB() *pg.DB {
	return postgresDB
}

func GetServerLocation() *time.Location {
	return serverLocation
}

func GetServerTime() time.Time {
	return time.Now().In(serverLocation)
}

func GetValidator() *validator.Validate {
	return validate
}

func GetSessionStore() *sessions.FilesystemStore {
	return sessionStore
}
