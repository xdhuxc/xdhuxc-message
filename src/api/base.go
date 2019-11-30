package api

import (
	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"

	"github.com/xdhuxc/xdhuxc-message/src/model"
	"github.com/xdhuxc/xdhuxc-message/src/service"
)

type BaseController struct {
	db   *gorm.DB
	bs   *service.BaseService
	ws   *restful.WebService
	conf *model.Configuration
}

func NewBaseController(conf *model.Configuration, db *gorm.DB) *BaseController {
	ws := new(restful.WebService)
	ws.Path("/api/v1").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	baseController := &BaseController{
		db:   db,
		bs:   service.NewBaseService(conf.DingTalkAuthentication, conf.EmailServer, db),
		ws:   ws,
		conf: conf,
	}

	// add new controller
	newHealthCheckController(baseController)
	newMessageController(baseController)
	newMetricsController(baseController)

	return baseController
}
