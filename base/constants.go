package base

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-pg/pg/v10"
)

var (
	idNode     *snowflake.Node
	postgresDB *pg.DB
)
