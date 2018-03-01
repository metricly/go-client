// +build !unit

package api

import (
	"github.com/metricly/go-client/model/core"
	"testing"
	"time"
)

func TestClientPostElement(t *testing.T) {
	//given
	client := NewClient("http://localhost:9400/ingest", "{api-key}")
	element := core.NewElement("elementId", "Element Name", "SERVER", "use1a")
	metric := core.NewMetric("metricId", "Metric Name", "bytes", "COUNTER", "None", core.Tag{"env", "prod"})
	sample, _ := core.NewSample("metricId", time.Now(), 0.0)
	element.AddAttribute("cpus", "4").AddTag("platform", "kubernetes").AddRelation("relatedElementId").AddMetric(metric).AddSample(sample)
	elements := []core.Element{element}
	//when & then
	if err := client.PostElements(elements); err != nil {
		t.Errorf("failed to post elements to test endpoint")
	}
}

func TestClientPostEvent(t *testing.T) {
	//given
	client := NewClient("http://localhost:9400/ingest", "{api-key}")
	event := core.NewEvent("metric go client", "test post event", "INFO", time.Now(), core.ElementMessage{"elementId", "INFO", "test"})
	event.AddTag("platform", "kubernetes")
	//when & then
	if err := client.PostEvents([]core.Event{event}); err != nil {
		t.Errorf("failed to post events to test endpoint")
	}
}

func TestClientPostCheck(t *testing.T) {
	//given
	client := NewClient("http://localhost:9400/ingest", "{api-key}")
	check := core.Check{"heartbeat", "elementId", 120}
	//when & then
	if err := client.PostCheck(check); err != nil {
		t.Errorf("failed to post check to test endpoint")
	}
}
