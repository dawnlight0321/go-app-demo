package main

import (
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/maxence-charriere/go-app/v6/pkg/log"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	version := os.Getenv("GAE_VERSION")

	log.Info("stating app engine server").
		T("addr", addr).
		T("version", version)

	h := &app.Handler{
		Title:   "Hello Google App Engine",
		Author:  "Maxence Charriere",
		Version: version,
	}

	if err := http.ListenAndServe(addr, h); err != nil {
		panic(err)
	}
}
