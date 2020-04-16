package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"

	log "github.com/sirupsen/logrus"
	_ "go.uber.org/automaxprocs"

	"github.com/xdhuxc/xdhuxc-message/src/api"
	"github.com/xdhuxc/xdhuxc-message/src/conf"
)

var cf = flag.String("conf", "conf.prod.yaml", "config path")

func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:8005", nil))
	}()

	flag.Parse()

	c, err := conf.InitConfiguration(*cf)
	if err != nil {
		log.Fatalf("init config error: %v", err)
		return
	}

	switch c.Log.Format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
	level, _ := log.ParseLevel(c.Log.Level)
	log.SetLevel(level)

	app := api.NewRouter()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
