package main

import (
	_ "git.rjbasitali.com/at-case-study/pkg/cache"
	_ "git.rjbasitali.com/at-case-study/pkg/db"
	_ "git.rjbasitali.com/at-case-study/pkg/log"
	"git.rjbasitali.com/at-case-study/server"
)

func main() {
	// blocking call
	server.Run()
}
