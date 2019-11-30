package api

import (
	"net/http"
	"path"

	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

func staticWs(c *restful.Container) {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/static/{subpath:*}").To(staticFromPathParam))

	c.Add(ws)
}

func staticFromPathParam(req *restful.Request, resp *restful.Response) {
	actual := path.Join("./static", req.PathParameter("subpath"))
	log.Errorf("serving %s ... (from %s)\n", actual, req.PathParameter("subpath"))

	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actual)
}
