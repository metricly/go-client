// +build !unit

package api

import (
	"testing"
	"time"
	"../model/core"
)

func TestClient(t *testing.T) {
	//given
	client := NewClient("http://localhost:9400/ingest", "43b6e3843e5db961fbc38cc24e796512")
	element := core.NewElement("elementId", "Element Name", "SERVER", "use1a")
	metric := core.NewMetric("metricId", "Metric Name", "bytes", "COUNTER", "None", core.Tag{"env", "prod"})
	sample, _ := core.NewSample("metricId", time.Now(), 0.0)
	element.AddAttribute("cpus", "4")
	element.AddTag("platform", "kubernetes")
	element.AddRelation("relatedElementId")
	element.AddMetric(metric)
	element.AddSample(sample)
	elements := []core.Element{element}
	//when
	err := client.post(elements)

	if err != nil {
		t.Errorf("failed to post elements to test endpoint")
	}
}
