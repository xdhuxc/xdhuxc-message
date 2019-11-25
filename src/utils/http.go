package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	"github.com/xdhuxc/xdhuxc-message/src/pkg"
)

// GET request
func DoGetWithURL(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, ioutil.NopCloser(bytes.NewBuffer([]byte(""))))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Infof("the request url is %s", url)
		log.Errorf("the response is not ok, the status code is %s", strconv.Itoa(resp.StatusCode))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorln(err)
		}
		log.Errorf("the response body is %s", string(body))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// POST request
func DoPostWithURL(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Infof("the request url is %s", url)
		log.Errorf("the response is not ok, the status code is %s", strconv.Itoa(resp.StatusCode))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorln(err)
		}
		log.Errorf("the response body is %s", string(body))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GET request
func DoGetWithRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Infof("the request url is %s", req.URL)
		log.Errorf("the response is not ok, the status code is %s", strconv.Itoa(resp.StatusCode))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorln(err)
		}
		log.Errorf("the response body is %s", string(body))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// 获取钉钉认证 token
func GetDingTalkToken(corporationID string, corporationSecret string) (string, error) {
	dingTalkTokenURL := fmt.Sprintf(pkg.DingTalkTokenURLTemplate, corporationID, corporationSecret)
	body, err := DoGetWithURL(dingTalkTokenURL)
	if err != nil {
		return "", err
	}

	errorCode := gjson.GetBytes(body, "errcode").Int()
	if errorCode != 0 {
		errorMessage := gjson.GetBytes(body, "errmsg").String()
		return "", errors.New(errorMessage)
	}

	return gjson.GetBytes(body, "access_token").String(), nil
}
