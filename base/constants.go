package base

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
)

var (
	idNode         *snowflake.Node
	postgresDB     *pg.DB
	serverLocation *time.Location
	validate       *validator.Validate
	sessionStore   *sessions.FilesystemStore
)
