package pkg

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"strconv"
	"strings"
)

type ResponseResult struct {
	Code int64       `json:"code"`
	Type string      `json:"type"`
	Data interface{} `json:"result"`
}

const (
	// 消息
	MessageHasBeenSent string = "200-0220001"

	// 参数错误
	ReadEntityError string = "400-0140001"

	// 权限错误
	AuthorityFailed   string = "401-0140003"
	InvalidParameters string = "400-0140004"

	// 参数缺失
	MissingMessageIDError     string = "400-0240001"
	MissingPageParameterError string = "400-0240002"

	// 消息错误
	MessageDataError         string = "400-0240003"
	HealthCheckError         string = "400-0240004"
	CreateMessageError       string = "400-0240005"
	SendMessageError         string = "400-0240006"
	UpdateMessageStatusError string = "400-0240007"
	NoSuchMessageError       string = "400-0240008"
	ListMessagesError        string = "400-0240009"
)

func WriteResponse(resp *restful.Response, code string, data interface{}) {
	httpCode, res := NewResponseResult(code, data)

	_ = resp.WriteHeaderAndEntity(httpCode, res)
}

func NewResponseResult(code string, result interface{}) (int, ResponseResult) {
	codes := strings.Split(code, "-")
	httpCode, _ := strconv.Atoi(codes[0])
	statusCode, _ := strconv.ParseInt(codes[1], 10, 64)

	return httpCode, ResponseResult{
		Code: statusCode,
		Type: "",
		Data: fmt.Sprint(result),
	}
}
