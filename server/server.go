package server

import (
	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/pkg/log"
)

func Run() {
	r := newRouter()
	if err := r.Run(cfg.APP_ADDR); err != nil {
		log.Error(err)
	}
}
