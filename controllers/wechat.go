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
	"programmer_wang/tools/wechat_message"
	"time"
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
	data := string(c.Ctx.Input.RequestBody)
	//c.Ctx.WriteString(string(c.Ctx.Input.RequestBody))

//	data := `
//		<xml>
//    <ToUserName><![CDATA[gh_bacfcce806e4]]></ToUserName>
//    <FromUserName><![CDATA[o9O7PwKec9bZRRejpiyFlAL_Sgwk]]></FromUserName>
//    <CreateTime>1541764082</CreateTime>
//    <MsgType><![CDATA[text]]></MsgType>
//    <Content><![CDATA[hhh]]></Content>
//    <MsgId>6621826310771202218</MsgId>
//    <Encrypt><![CDATA[n/ZdEQL/mcsR/ZeQVjuyqeX9fnaqWbS9zKTAt5b9+/TCv3Vo/l882Cjruji5gcQfsz3SJvLT/7+UT7rwu/neqIstGxox7pPYjKF6GFi3EZXyL+hQHuNH5kATiy6Xe/cH/Mqko8BdhKdB9RdT/NcBrGT1oWqkOeWNOLlfzVhHZxQZd3Yc2pvAg1UnLZxm1DX+V6Sfqud/s+kqLlzbpjc4v3O3EOhx/ja8ZtAXK4LipJ2Mt4aa3fjxpxYy1JSb1Dn+st1K2alw/qBX4TenvQpn+GBdDrCWnLDpSAM3ZDL+HtUDj4zAK5Tlun8YDJQKdQg884tbqVQgAvG2dVhyY4UAYacFZxP3+AhCrrA3CkiaqVPIjSMBGZVMsM9B3mNfql/hYKsx4ELhwN27PUdpNm+2s9+vj1MGl7tqXB8QGXIaj3A=]]></Encrypt>
//</xml>
//	`
	info, _ := passive_response.DecodeTextInfo(data, &passive_response.TextInfo{})
	log.Println("here is info")
	log.Println(info)
	//c.Ctx.WriteString(info.Content)
	//c.Ctx.WriteString(string(c.Ctx.Input.RequestBody))

	x := passive_response.TextInfo{}
	x.FromUserName = (*info).ToUserName
	x.ToUserName = (*info).FromUserName
	x.CreateTime = string(time.Now().Unix())
	x.MsgType = "text"
	x.Content = fmt.Sprintf("hello %s", (*info).FromUserName)
	c.Data["xml"] = &x

	c.ServeXML()
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