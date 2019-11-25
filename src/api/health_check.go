package api

import (
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"

	"github.com/xdhuxc/xdhuxc-message/src/model"
	"github.com/xdhuxc/xdhuxc-message/src/pkg"
)

type healthCheckController struct {
	*BaseController
}

func newHealthCheckController(bc *BaseController) *healthCheckController {
	tags := []string{"xdhuxc-message-health_check"}
	hcc := &healthCheckController{bc}

	hcc.ws.Route(hcc.ws.GET("/healthcheck").
		To(hcc.Get).
		Doc("health check").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, "OK", model.Result{}).
		Returns(http.StatusBadRequest, "ERROR", model.Result{}))

	return hcc
}

func (hcc *healthCheckController) Get(req *restful.Request, resp *restful.Response) {
	result, err := hcc.bs.HealthCheckService.Get()
	if err != nil {
		pkg.WriteResponse(resp, pkg.HealthCheckError, err)
		return
	}

	_ = resp.WriteEntity(model.NewResult(0, nil, result))
}
