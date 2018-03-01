package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/metricly/go-client/model/core"
	"net/http"
	"strconv"
	"strings"
)

//Client is a Rest API Client for ingesting Elements, Events and Checks to Metricly Cloud, use its Construction function NewClient to create a Client
type Client struct {
	elementIngestEndpoint string
	eventIngestEndpoint   string
	checkIngestEndpoint   string
	client                *http.Client
}

//NewClient constructs a Client given its ingest URL and apiKey
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

//PostElement posts a slice of Elements to cloud ingest API
func (c *Client) PostElements(elements []core.Element) error {
	payload, _ := json.Marshal(elements)
	req, _ := http.NewRequest("POST", c.elementIngestEndpoint, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return postPayload(c, req, "element")
}

//PostCheck posts a check to cloud ingest API
func (c *Client) PostCheck(check core.Check) error {
	checkUrl := strings.Join([]string{c.checkIngestEndpoint, check.Name, check.ElementId, strconv.Itoa(check.TTL)}, "/")
	req, _ := http.NewRequest("POST", checkUrl, bytes.NewBuffer([]byte{}))
	return postPayload(c, req, "check")
}

//PostEvent posts a slice of Events to cloud ingest API
func (c *Client) PostEvents(events []core.Event) error {
	payload, _ := json.Marshal(events)
	req, _ := http.NewRequest("POST", c.eventIngestEndpoint, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return postPayload(c, req, "event")
}

func postPayload(c *Client, request *http.Request, entity string) error {
	resp, err := c.client.Do(request)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusAccepted {
		result, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%s payload is not accepted due to response status: %d body: %d", entity, resp.StatusCode, string(result))
	}
	defer resp.Body.Close()
	return nil
}
