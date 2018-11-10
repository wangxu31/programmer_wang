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

func DecodeTextInfo(xmlString string, info *TextInfo) (*TextInfo, error) {
	err := xml.Unmarshal([]byte(xmlString), &info)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	log.Println(info)
	return info, nil
}