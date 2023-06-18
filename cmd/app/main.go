package main

import (
	"git.rjbasitali.com/at-case-study/pkg/cache"
	"git.rjbasitali.com/at-case-study/pkg/db"
	"git.rjbasitali.com/at-case-study/server"
)

func main() {
	db.Init()    // initialize the database
	cache.Init() // initialize redis client
	server.Run() // run the server - blocking call
}
