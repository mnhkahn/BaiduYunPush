package BaiduYunPush

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type BaiduYunPush struct {
	Seckey  string
	UrlBase string
	Method  string
	query   map[string]string
}

func New(api_key string, seckey string) *BaiduYunPush {
	push := new(BaiduYunPush)
	push.query = map[string]string{}

	push.Seckey = seckey
	push.UrlBase = "http://channel.api.duapp.com/rest/2.0/channel/channel"
	push.Method = "POST"
	push.query["apikey"] = api_key
	push.query["message_type"] = "1"
	push.query["messages"] = ""
	push.query["method"] = "push_msg"
	push.query["msg_keys"] = "msgkey"
	push.query["push_type"] = "3"
	push.query["timestamp"] = ""

	return push
}

func (push *BaiduYunPush) Push(title string, description string) (bool, error) {
	push.query["messages"] = "{\"title\":\"" + title + "\",\"description\":\"" + description + "\"}"
	push.query["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)

	sign := push.Method + push.UrlBase
	for k, v := range push.query {
		sign += k + "=" + v
	}

	sign += push.Seckey

	sign = url.QueryEscape(sign)

	m := md5.New()
	m.Write([]byte(sign))
	sign = hex.EncodeToString(m.Sum(nil))

	url_sign := push.UrlBase + "?"
	for k, v := range push.query {
		url_sign += k + "=" + v + "&"
	}

	url_sign += "sign=" + sign

	resp, err := HTTPPost(url_sign, nil, nil)
	if err != nil {
		e := BaiduResponseError{}
		json.Unmarshal([]byte(resp), &e)
		return false, fmt.Errorf("id:%s, %s", e.RequestId, e.ErrorMessage)
	} else {
		return true, nil
	}

}
