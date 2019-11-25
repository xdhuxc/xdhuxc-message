package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSON_Scan(t *testing.T) {
	as := []string{"a", "b", "c"}
	dataInBytes, err := json.Marshal(as)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(dataInBytes))
	}
}
