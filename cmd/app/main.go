package main

import (
	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/pkg/cache"
	"git.rjbasitali.com/at-case-study/pkg/db"
	"git.rjbasitali.com/at-case-study/server"
)

func main() {
	// initialize the config - should be done before anything else
	cfg.Init()

	// initialize the database
	db.Init(cfg.DB_CONN_STR)

	// initialize redis client
	cache.Init(cfg.REDIS_HOST, cfg.REDIS_PORT, cfg.REDIS_PASSWORD, cfg.REDIS_DB)

	// run the server - blocking call
	server.Run()
}
