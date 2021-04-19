package base

import (
	"context"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
)

var (
	idNode         *snowflake.Node
	postgresDB     *pg.DB
	redisDB        *redis.Client
	serverLocation *time.Location
	validate       *validator.Validate
	sessionStore   *sessions.FilesystemStore
	ctx            context.Context = context.TODO()
)
