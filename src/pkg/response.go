package pkg

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"strconv"
	"strings"
)

type ResponseResult struct {
	Code   int64       `json:"code"`
	Type   string      `json:"type"`
	Result interface{} `json:"result"`
}

const (
	ReadEntityError  string = "400-10001"
	MessageDataError string = "400-10002"

	AuthFailed        string = "401-10003"
	InvalidParameters string = "400-10004"

	HealthCheckError         string = "500-100005"
	CreateMessageError       string = "500-100006"
	SendMessageError         string = "500-100007"
	UpdateMessageStatusError string = "500-100008"
	NoSuchMessageError       string = "500-100009"
	MessageHasBeenSentError  string = "500-100010"
	AbsentOfMessageIDError   string = "500-100010"
	ListMessagesError        string = "500-100010"
)

func WriteResponse(resp *restful.Response, code string, result interface{}) {
	httpCode, res := NewResponseResult(code, result)

	_ = resp.WriteHeaderAndEntity(httpCode, res)
}

func NewResponseResult(code string, result interface{}) (int, ResponseResult) {
	codes := strings.Split(code, "-")
	httpCode, _ := strconv.Atoi(codes[0])
	statusCode, _ := strconv.ParseInt(codes[1], 10, 64)

	return httpCode, ResponseResult{
		Code:   statusCode,
		Type:   "",
		Result: fmt.Sprint(result),
	}
}
