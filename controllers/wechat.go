package controllers

import (
	"github.com/astaxie/beego"
	"sort"
	"strings"
	"log"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"encoding/json"
	"programmer_wang/tools/request"
)

const TOKEN = "clive31"
const APP_ID = "1wxf344082d04f03fbc"
const APP_SECRET = "3a43d7b3e3d6626d8274e2da6c8e327c"

type ResponseStruct struct {
	Errcode int
	Errmsg string
	Access_token string
	Expires_in int
}

type WeChatController struct {
	beego.Controller
}

func (c *WeChatController) Verify() {
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")
	token := TOKEN
	data := []string{timestamp, nonce, token}
	sort.Strings(data)

	res := strings.Join(data, "")

	resSha1 := sha1s(res)
	log.Println("resSha1", resSha1)
	log.Println("signature", signature)
	log.Println("echostr", echostr)

	if resSha1 != signature {
		log.Println("signature not equal")
		c.Ctx.WriteString("")
	}

	c.Ctx.WriteString(echostr)
}

func (c *WeChatController) Answer() {
	log.Println(c.Ctx.Input.RequestBody)
	log.Println(string(c.Ctx.Input.RequestBody))
	c.Ctx.WriteString(string(c.Ctx.Input.RequestBody))
}

func sha1s(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func GetAccessToken() string {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", APP_ID, APP_SECRET)
	response, _ := httpRequest.DoHttpGet(url)

	resStruct := ResponseStruct{}
	// var f interface{}
	// json.Unmarshal(b, &f)

	err := json.Unmarshal([]byte(response), &resStruct)
	if err != nil {
		log.Println(err)
		return ""
	}

	if resStruct.Errcode != 0 {
		log.Println(resStruct.Errmsg)
		return ""
	}

	accessToken := resStruct.Access_token
	return accessToken
	//15_sbuM1R1cQaZiTh3xWB_Es-8Mf4Cc6_iGAkT-4doJI_4owe0lgIUvtjqd_liZNHvVckym2DchtaePxvpBeO_pJGdcHqV9ZKG4K3oIc9d-Jr4yOAQNXOWyl-sNuRgCRQjAJAFUT
}