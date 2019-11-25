package api

import (
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/xdhuxc/xdhuxc-message/src/model"
	"github.com/xdhuxc/xdhuxc-message/src/pkg"
	"net/http"
	"strconv"
	"time"
)

type MessageController struct {
	*BaseController
}

func newMessageController(baseController *BaseController) *MessageController {
	tags := []string{"xdhuxc-message_message"}
	mc := &MessageController{baseController}

	mc.ws.Route(mc.ws.POST("/message").
		To(mc.Create).
		Reads(model.Message{}).
		Doc("receive a message").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, "OK", model.Result{}).
		Returns(http.StatusBadRequest, "ERROR", model.Result{}))

	mc.ws.Route(mc.ws.POST("/message/{id}/retry").
		To(mc.Retry).
		Doc("retry to send the specified message if it has not been sent").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, "OK", model.Result{}).
		Returns(http.StatusBadRequest, "ERROR", model.Result{}))

	mc.ws.Route(mc.ws.POST("/message/{id}/again").
		To(mc.Again).
		Doc("send the specified message again, whether it has been sent or not").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(http.StatusOK, "OK", model.Result{}).
		Returns(http.StatusBadRequest, "ERROR", model.Result{}))

	return mc
}

func (mc *MessageController) Create(req *restful.Request, resp *restful.Response) {
	var m model.Message
	var result *model.Message
	var cme error

	if err := req.ReadEntity(&m); err != nil {
		pkg.WriteResponse(resp, pkg.ReadEntityError, err)
		return
	}
	// 校验 Message 的合法性
	if err := m.Validate(); err != nil {
		pkg.WriteResponse(resp, pkg.MessageDataError, err)
		return
	}

	// 发送消息
	if m.MessageType == pkg.MessageTypeDingTalk {
		m.User = m.Sender
		sme := mc.bs.DingTalkService.Send(*m.ConvertToDingTalk())
		if sme != nil {
			m.IsSent = false
			result, cme = mc.bs.MessageService.Create(m)
			if cme != nil {
				pkg.WriteResponse(resp, pkg.CreateMessageError, cme)
				return
			}
			// 操作审计
			_ = mc.bs.AuditService.Create(model.OperationAudit{
				User:       m.User,
				Operate:    model.OperatingTypeAdd,
				Object:     model.MessageTypeDingTalk,
				CreateTime: time.Now(),
			})
			pkg.WriteResponse(resp, pkg.SendMessageError, sme)
			return
		}
	} else if m.MessageType == pkg.MessageTypeEmail {
		m.User = mc.conf.EmailServer.User
		sme := mc.bs.EmailService.Send(*m.ConvertToEmail())
		if sme != nil {
			m.IsSent = false
			result, cme = mc.bs.MessageService.Create(m)
			if cme != nil {
				pkg.WriteResponse(resp, pkg.CreateMessageError, cme)
				return
			}
			// 操作审计
			_ = mc.bs.AuditService.Create(model.OperationAudit{
				User:       m.User,
				Operate:    model.OperatingTypeAdd,
				Object:     model.MessageTypeWeChat,
				CreateTime: time.Now(),
			})
			pkg.WriteResponse(resp, pkg.SendMessageError, sme)
			return
		}
	}

	// 入库
	m.IsSent = true
	result, err := mc.bs.MessageService.Create(m)
	if err != nil {
		pkg.WriteResponse(resp, pkg.CreateMessageError, err)
		return
	}

	_ = mc.bs.AuditService.Create(model.OperationAudit{
		User:       m.User,
		Operate:    model.OperatingTypeAdd,
		Object:     m.MessageType,
		CreateTime: time.Now(),
	})

	_ = resp.WriteEntity(model.NewResult(0, nil, result))
}

func (mc *MessageController) Retry(req *restful.Request, resp *restful.Response) {
	messageID, err := strconv.ParseInt(req.PathParameter("id"), 10, 64)
	m, err := mc.bs.MessageService.GetMessageByID(messageID)
	if err != nil {
		pkg.WriteResponse(resp, pkg.NoSuchMessageError, err)
		return
	}
	if m.IsSent {
		pkg.WriteResponse(resp, pkg.MessageHasBeenSentError, err)
		return
	}

	// 发送消息
	if m.MessageType == pkg.MessageTypeDingTalk {
		err = mc.bs.DingTalkService.Send(*m.ConvertToDingTalk())
		if err != nil {
			pkg.WriteResponse(resp, pkg.SendMessageError, err)
			return
		}
	} else if m.MessageType == pkg.MessageTypeEmail {
		err = mc.bs.EmailService.Send(*m.ConvertToEmail())
		if err != nil {
			pkg.WriteResponse(resp, pkg.SendMessageError, err)
			return
		}
	}
	// 状态调整
	err = mc.bs.MessageService.UpdateStatus(*m)
	if err != nil {
		pkg.WriteResponse(resp, pkg.UpdateMessageStatusError, err)
		return
	}

	_ = resp.WriteEntity(model.NewResult(0, nil, m))
}

func (mc *MessageController) Again(req *restful.Request, resp *restful.Response) {
	messageID, err := strconv.ParseInt(req.PathParameter("id"), 10, 64)
	m, err := mc.bs.MessageService.GetMessageByID(messageID)
	if err != nil {
		pkg.WriteResponse(resp, pkg.NoSuchMessageError, err)
		return
	}

	// 发送消息
	if m.MessageType == pkg.MessageTypeDingTalk {
		err = mc.bs.DingTalkService.Send(*m.ConvertToDingTalk())
		if err != nil {
			pkg.WriteResponse(resp, pkg.SendMessageError, err)
			return
		}
	} else if m.MessageType == pkg.MessageTypeEmail {
		err = mc.bs.EmailService.Send(*m.ConvertToEmail())
		if err != nil {
			pkg.WriteResponse(resp, pkg.SendMessageError, err)
			return
		}
	}
	// 状态调整
	if !m.IsSent {
		err = mc.bs.MessageService.UpdateStatus(*m)
		if err != nil {
			pkg.WriteResponse(resp, pkg.UpdateMessageStatusError, err)
			return
		}
	}

	_ = resp.WriteEntity(model.NewResult(0, nil, m))
}
