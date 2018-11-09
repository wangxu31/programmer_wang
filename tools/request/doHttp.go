package httpRequest

import (
	"net/http"
	"log"
	"io/ioutil"
)

func DoHttpGet(url string) (accessToken string, err error) {
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", url, nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		log.Println(parseFormErr)
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