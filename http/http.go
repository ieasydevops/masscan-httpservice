package http

import (
	"github.com/meixinyun/masscan-httpservice/g"
	"log"
	"net/http"
)

func init() {
	configHealthRoutes()
}

func Start() {

	if !g.Config().Http.Enabled {
		return
	}

	addr := g.Config().Http.Listen
	if addr == "" {
		return
	}

	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}

	log.Println("listening", addr)
	log.Fatalln(s.ListenAndServe())
}
