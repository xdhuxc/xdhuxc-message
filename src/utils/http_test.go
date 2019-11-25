package utils

import "testing"

func TestGetDingTalkToken(t *testing.T) {
	corporationID := ""
	corporationSecret := ""

	token, err := GetDingTalkToken(corporationID, corporationSecret)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
}
