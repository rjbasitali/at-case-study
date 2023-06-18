package main

import (
	"git.rjbasitali.com/at-case-study/pkg/cache"
	"git.rjbasitali.com/at-case-study/pkg/db"
	_ "git.rjbasitali.com/at-case-study/pkg/log"
	"git.rjbasitali.com/at-case-study/server"
)

func main() {
	// initialize the database
	db.Init()
	// initialize redis client
	cache.Init()
	// run the server - blocking call
	server.Run()
}
