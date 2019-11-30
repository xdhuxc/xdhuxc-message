package model

import "encoding/json"

type MetricsResult struct {
	TotalMessages          int64 `json:"totalMessages"`
	TotalFailedMessages    int64 `json:"totalFailedMessages"`
	TotalSentMessages      int64 `json:"totalSuccessfulMessages"`
	DingTalkMessages       int64 `json:"dingTalkMessages"`
	DingTalkFailedMessages int64 `json:"dingTalkFailedMessages"`
	DingTalkSentMessages   int64 `json:"DingTalkSentMessages"`
	EmailMessages          int64 `json:"emailMessages"`
	EmailFailedMessages    int64 `json:"emailFailedMessages"`
	EmailSentMessages      int64 `json:"emailSentMessages"`
}

func (mr *MetricsResult) String() string {
	if dataInBytes, err := json.Marshal(&mr); err == nil {
		return string(dataInBytes)
	}

	return ""
}
