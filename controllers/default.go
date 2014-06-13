package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"net/url"
	"strconv"
	"time"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

var apikey = "vkl8Pc6QCUjHSemG0wUwVAzQ"
var seckey = "GbrgjnajK6LKl5oa5EIGCuGBRYe8twRP"
var method = "POST"
var url_base = "http://channel.api.duapp.com/rest/2.0/channel/channel"
var url_base1 = "channel.api.duapp.com/rest/2.0/channel/channel"
var query = map[string]string{}

func (this *MainController) Push() {
	title := this.GetString("title")
	description := this.GetString("description")

	query["apikey"] = apikey
	query["message_type"] = "1"
	query["messages"] = "{\"title\":\"" + title + "\",\"description\":\"" + description + "\"}"
	query["method"] = "push_msg"
	query["msg_keys"] = "msgkey"
	query["push_type"] = "3"
	query["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)

	sign := method + url_base
	for k, v := range query {
		sign += k + "=" + v
	}

	sign += seckey

	sign = url.QueryEscape(sign)

	m := md5.New()
	m.Write([]byte(sign))
	sign = hex.EncodeToString(m.Sum(nil))

	url_sign := url_base1 + "?"
	for k, v := range query {
		url_sign += k + "=" + v + "&"
	}

	url_sign += "&sign=" + sign

	req := httplib.Post(url_sign)
	resp, _ := req.Response()

	this.Ctx.ResponseWriter.WriteHeader(resp.StatusCode)
}
