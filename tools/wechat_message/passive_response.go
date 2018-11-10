package passive_response

import (
	"encoding/xml"
	"log"
)

type TextInfo struct {
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId"`
	Encrypt      string   `xml:"Encrypt"`
}

const TextReplyFormat = `<xml>
    <ToUserName><![CDATA[gh_bacfcce806e4]]></ToUserName>
    <FromUserName><![CDATA[o9O7PwKec9bZRRejpiyFlAL_Sgwk]]></FromUserName>
    <CreateTime>1541764082</CreateTime>
    <MsgType><![CDATA[text]]></MsgType>
    <Content><![CDATA[hhh]]></Content>
    <MsgId>6621826310771202218</MsgId>
    <Encrypt><![CDATA[n/ZdEQL/mcsR/ZeQVjuyqeX9fnaqWbS9zKTAt5b9+/TCv3Vo/l882Cjruji5gcQfsz3SJvLT/7+UT7rwu/neqIstGxox7pPYjKF6GFi3EZXyL+hQHuNH5kATiy6Xe/cH/Mqko8BdhKdB9RdT/NcBrGT1oWqkOeWNOLlfzVhHZxQZd3Yc2pvAg1UnLZxm1DX+V6Sfqud/s+kqLlzbpjc4v3O3EOhx/ja8ZtAXK4LipJ2Mt4aa3fjxpxYy1JSb1Dn+st1K2alw/qBX4TenvQpn+GBdDrCWnLDpSAM3ZDL+HtUDj4zAK5Tlun8YDJQKdQg884tbqVQgAvG2dVhyY4UAYacFZxP3+AhCrrA3CkiaqVPIjSMBGZVMsM9B3mNfql/hYKsx4ELhwN27PUdpNm+2s9+vj1MGl7tqXB8QGXIaj3A=]]></Encrypt>
</xml>
`

func DecodeTextInfo(xmlString string, info *TextInfo) (*TextInfo, error) {
	err := xml.Unmarshal([]byte(xmlString), &info)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	log.Println(info)
	return info, nil
}