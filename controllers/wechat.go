package controllers

import (
	"github.com/astaxie/beego"
	"sort"
	"log"
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
	log.Println(echostr)
	c.ServeJSON()
}
