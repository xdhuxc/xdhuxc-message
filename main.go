package main

import (
	"flag"
	"os"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/xdhuxc/xdhuxc-message/src/api"
	"github.com/xdhuxc/xdhuxc-message/src/conf"
)

var cf = flag.String("conf", "conf.prod.yaml", "config path")

func init() {
	cpus := os.Getenv("CPUS")
	if nums, err := strconv.Atoi(cpus); err != nil {
		runtime.GOMAXPROCS(nums)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

func main() {
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
