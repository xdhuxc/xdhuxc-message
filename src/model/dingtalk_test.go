package model

import "testing"

func TestDingTalk_ConvertToDingTalkMessage(t *testing.T) {
	receivers := []string{"a", "b", "c"}
	var users string
	for index, r := range receivers {
		if index == 0 {
			users = r
		} else {
			users = users + "," + r
		}
	}

	if users == "a,b,c" {
		t.Log(users)
	} else {
		t.Fail()
	}
}
