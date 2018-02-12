// +build !unit

package api

import (
	"testing"
	"time"
	"../model/core"
	"log"
)

func TestClientPostElement(t *testing.T) {
	//given
	client := NewClient("http://localhost:9400/ingest", "43b6e3843e5db961fbc38cc24e796512")
	element := core.NewElement("elementId", "Element Name", "SERVER", "use1a")
	metric := core.NewMetric("metricId", "Metric Name", "bytes", "COUNTER", "None", core.Tag{"env", "prod"})
	sample, _ := core.NewSample("metricId", time.Now(), 0.0)
	element.AddAttribute("cpus", "4").AddTag("platform", "kubernetes").AddRelation("relatedElementId").AddMetric(metric).AddSample(sample)
	elements := []core.Element{element}
	//when & then
	if err := client.PostElement(elements); err != nil {
		log.Println("failed to post elements to test endpoint with error: ", err)
		t.Errorf("failed to post elements to test endpoint")
	}
}

func TestClientPostEvent(t *testing.T) {
	//given
	client := NewClient("http://localhost:9400/ingest", "43b6e3843e5db961fbc38cc24e796512")
	event := core.NewEvent("metric go client", "test post event", "INFO", time.Now(), core.ElementMessage{"elementId", "INFO", "test"})
	event.AddTag("platform", "kubernetes")
	//when & then
	if err := client.PostEvent([]core.Event{event}); err != nil {
		log.Println("failed to post events to test endpoint with error: ", err)
		t.Errorf("failed to post events to test endpoint")
	}
}

func TestClientPostCheck(t *testing.T) {
	//given
	client := NewClient("http://localhost:9400/ingest", "43b6e3843e5db961fbc38cc24e796512")
	check := core.Check{"heartbeat", "elementId", 120}
	//when & then
	if err := client.PostCheck(check); err != nil {
		log.Println("failed to post check to test endpoint with error: ", err)
		t.Errorf("failed to post check to test endpoint")
	}
}
