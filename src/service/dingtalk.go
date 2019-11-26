package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"

	"github.com/xdhuxc/xdhuxc-message/src/model"
	"github.com/xdhuxc/xdhuxc-message/src/pkg"
	"github.com/xdhuxc/xdhuxc-message/src/utils"
)

type dingTalkService struct {
	dingTalkAuthentication model.DingTalkAuthentication
}

func newDingTalkService(dta model.DingTalkAuthentication) *dingTalkService {
	return &dingTalkService{dta}
}

func (dts *dingTalkService) ConvertToDingTalkMessage(dt model.DingTalk) ([]byte, error) {
	var receivers []string
	err := json.Unmarshal(dt.Receivers, &receivers)
	if err != nil {
		return nil, err
	}

	var userIDList string
	for index, r := range receivers {
		if index == 0 {
			userIDList = r
		} else {
			userIDList = userIDList + "," + r
		}
	}

	dtm := &model.DingTalkMessage{
		AgentID:    dts.dingTalkAuthentication.AgentID,
		UserIDList: userIDList,
		DingTalkMessageBody: model.DingtalkMessageBody{
			DingTalkMessageType: "text",
			DingTalkMessageText: model.DingTalkMessageText{
				Content: dt.Message,
			},
		},
	}

	return json.Marshal(&dtm)
}

func (dts *dingTalkService) Send(dt model.DingTalk) error {
	token, err := utils.GetDingTalkToken(dts.dingTalkAuthentication.CorporationID, dts.dingTalkAuthentication.CorporationSecret)
	if err != nil {
		return err
	}
	messageURL := fmt.Sprintf(pkg.DingTalkSendMessageURLTemplate, token)
	data, err := dts.ConvertToDingTalkMessage(dt)
	if err != nil {
		return err
	}

	body, err := utils.DoPostWithURL(messageURL, data)
	if err != nil {
		return err
	}
	errorCode := gjson.GetBytes(body, "errcode").Int()
	if errorCode != 0 {
		errorMessage := gjson.GetBytes(body, "errmsg").String()
		return errors.New(errorMessage)
	}

	return nil
}
