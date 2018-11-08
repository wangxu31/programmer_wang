package controllers

import (
	"github.com/astaxie/beego"
	"sort"
	"strings"
	"log"
	"crypto/sha1"
	"encoding/hex"
)

const TOKEN = "clive31"

type WeChatController struct {
	beego.Controller
}

func (c *WeChatController) Verify() {
	//signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")
	token := TOKEN
	data := []string{timestamp, nonce, token}
	sort.Strings(data)

	res := strings.Join(data, "")

	resSha1 := sha1s(res)
	log.Println("resSha1", resSha1)
	log.Println("echostr", echostr)

	c.Ctx.WriteString(echostr)
}

func sha1s(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}
