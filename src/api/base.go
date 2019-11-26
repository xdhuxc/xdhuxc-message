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
	conf model.Configuration
	mws  *restful.WebService
}

func NewBaseController(conf model.Configuration, db *gorm.DB) *BaseController {
	ws := new(restful.WebService)

	ws.Path("/api/v1").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	mws := new(restful.WebService)
	mws.Path("/metrics").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	baseController := &BaseController{
		db:   db,
		bs:   service.NewBaseService(conf.DingTalkAuthentication, conf.EmailServer, db),
		ws:   ws,
		conf: conf,
		mws:  mws,
	}

	// add new controller
	newHealthCheckController(baseController)
	newMessageController(baseController)

	return baseController
}
