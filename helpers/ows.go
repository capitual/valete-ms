package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/capitual/valete_ms/config"
)

func OwsRequest(method string, url string, body string) (map[string]interface{}, error) {
	var jsonStr = []byte(body)
	req, err := http.NewRequest(method, config.GetConfig().OwsUrlBase+url, bytes.NewBuffer(jsonStr))

	if err != nil {
		return nil, err
	}

	nonce := strconv.Itoa(int(time.Now().UnixNano()))

	signature := SignRequest("/v1/"+url, nonce, body)
	req.Header.Add("X-API-Key", config.GetConfig().OwsKey)
	req.Header.Add("X-Nonce", nonce)
	req.Header.Add("X-Signature", signature)

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(responseBody)), &jsonMap)

	return jsonMap, nil
}
