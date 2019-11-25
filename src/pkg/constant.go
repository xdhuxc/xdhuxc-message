package pkg

const (
	MessageTypeEmail    = "email"
	MessageTypeDingTalk = "dingtalk"
	MessageTypeWeChat   = "wechat"
)

const (
	DingTalkTokenURLTemplate       = "https://oapi.dingtalk.com/gettoken?corpid=%s&corpsecret=%s"
	DingTalkSendMessageURLTemplate = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s"
)
