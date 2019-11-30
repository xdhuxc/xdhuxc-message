package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"

	"github.com/xdhuxc/xdhuxc-message/src/conf"
	"github.com/xdhuxc/xdhuxc-message/src/database"
)

type Router struct {
	container *restful.Container
	bs        *BaseController
}

func NewRouter() *Router {
	mysqldb, err := database.NewMysqlClient(conf.GetConfiguration().Database)
	if err != nil {
		log.Fatalf("new mysql client error: %v\n", err)
		return nil
	}

	baseController := NewBaseController(conf.GetConfiguration(), mysqldb)
	container := restful.NewContainer()
	container.Add(baseController.ws)

	metrics(container, baseController)
	swagger(container, conf.GetConfiguration().Address)
	staticWs(container)

	baseController.ws.Filter(baseController.metrics)
	baseController.ws.Filter(baseController.Page)

	r := &Router{
		container: container,
		bs:        baseController,
	}

	return r
}

func (r *Router) Run() error {
	log.Infof("start http server: %s", conf.GetConfiguration().Address)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", conf.GetConfiguration().Address),
		Handler:      r.container,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}
