package controllers

import (
	"github.com/astaxie/beego"
	"sort"
)

const TOKEN = "clive31"

type WeChatController struct {
	beego.Controller
}

func (c *WeChatController) Verify() {
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")
	token := TOKEN
	data := []string{signature, timestamp, nonce, echostr, token}
	sort.Strings(data)
	c.Data["json"] = echostr
	c.ServeJSON()
}
