package main

import (
	"flag"
	"fmt"
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

	err := conf.InitConfiguration(*cf)
	if err != nil {
		log.Fatalf("init config error: %v", err)
		return
	}
	app := api.NewRouter()

	if err := app.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		log.Fatal(err)
	}

}
