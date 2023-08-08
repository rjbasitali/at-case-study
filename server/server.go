package server

import (
	"fmt"

	"git.rjbasitali.com/at-case-study/cfg"
	"github.com/gin-gonic/gin"
)

// Run runs the server.
// It panics if the server could not be started.
func Run() {
	r := newRouter()
	if err := r.Run(cfg.APP_ADDR); err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, "error running server: ", err)
	}
}
