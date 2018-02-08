package api

import (
	"../model/core"
	"net/http"
	"strings"
	"encoding/json"
	"bytes"
	"log"
	"io/ioutil"
)

type Client struct {
	endpoint string
	client *http.Client
}

func NewClient(url, apiKey string) Client {

	return Client{
		strings.TrimSuffix(url, "/") + "/" + apiKey,
		&http.Client{},
	}
}

func (c *Client) post(elements []core.Element) error {
	payload, _ := json.Marshal(elements)
	req, _ := http.NewRequest("POST", c.endpoint, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := c.client.Do(req)
	log.Println("posting payload to api intest endpoint: ", c.endpoint, " with payload: ", string(payload), " got response status: ", resp.StatusCode)
	if resp.StatusCode > 399 {
		result, _ := ioutil.ReadAll(resp.Body)
		log.Println("response error: " + string(result))
	}
	if err != nil {
		log.Println("error posting payload to api ingest endpoint: ", c.endpoint, err)
	}
	defer resp.Body.Close()
	return nil
}