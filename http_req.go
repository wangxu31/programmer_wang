package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	"errors"
)

const APP_ID = "1wxf344082d04f03fbc"
const APP_SECRET = "3a43d7b3e3d6626d8274e2da6c8e327c"

type ResponseStruct struct {
	Errcode int
	Errmsg string
	Access_token string
	Expires_in int
}

func main() {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", APP_ID, APP_SECRET)
	response, _ := DoHttpGet(url)

	resStruct := ResponseStruct{}
	// var f interface{}
	// json.Unmarshal(b, &f)

	err := json.Unmarshal([]byte(response), &resStruct)
	if err != nil {
		log.Println(err)
		return
	}

	if resStruct.Errcode != 0 {
		log.Println(resStruct.Errmsg)
		return
	}

	accessToken := resStruct.Access_token
	fmt.Println(accessToken)
	//15_sbuM1R1cQaZiTh3xWB_Es-8Mf4Cc6_iGAkT-4doJI_4owe0lgIUvtjqd_liZNHvVckym2DchtaePxvpBeO_pJGdcHqV9ZKG4K3oIc9d-Jr4yOAQNXOWyl-sNuRgCRQjAJAFUT
}

func DoHttpGet(url string) (accessToken string, err error) {
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", url, nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	if resp.Status != "200 OK" {
		log.Println(string(respBody))
		return "", errors.New("request failed")
	}

	log.Println("response Body : ", string(respBody))
	return string(respBody), nil
}


