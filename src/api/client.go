package api

import (
	"../model/core"
	"net/http"
	"strings"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"fmt"
	"strconv"
)

type Client struct {
	elementIngestEndpoint string
	eventIngestEndpoint string
	checkIngestEndpoint string
	client *http.Client
}

func NewClient(url, apiKey string) Client {
	elementIngestEndpoint := strings.TrimSuffix(url, "/") + "/" + apiKey
	eventIngestEndpoint := strings.Replace(elementIngestEndpoint, "/ingest/", "/ingest/events/", 1)
	checkIngestEndpoint := strings.Replace(elementIngestEndpoint, "/ingest/", "/check/", 1)
	return Client{
		elementIngestEndpoint,
		eventIngestEndpoint,
		checkIngestEndpoint,
		&http.Client{},
	}
}

func (c *Client) PostElement(elements []core.Element) error {
	payload, _ := json.Marshal(elements)
	req, _ := http.NewRequest("POST", c.elementIngestEndpoint, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return postPayload(c, req, "element")
}

func (c *Client) PostCheck(check core.Check) error {
	checkUrl := strings.Join([]string{c.checkIngestEndpoint, check.Name, check.ElementId, strconv.Itoa(check.TTL)}, "/")
	req, _ := http.NewRequest("POST", checkUrl, bytes.NewBuffer([]byte{}))
	return postPayload(c, req, "check")
}

func (c *Client) PostEvent(events []core.Event) error {
	payload, _ := json.Marshal(events)
	req, _ := http.NewRequest("POST", c.eventIngestEndpoint, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return postPayload(c, req, "event")
}

func postPayload(c *Client, request *http.Request, entity string) error {
	resp, err := c.client.Do(request);
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusAccepted {
		result, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%s payload is not accepted due to response status: %d body: %d", entity, resp.StatusCode, string(result))
	}
	defer resp.Body.Close()
	return nil
}