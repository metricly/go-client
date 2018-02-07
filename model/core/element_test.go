package core

import (
	"testing"
	"time"
	"log"
)

func TestAddElementSampleAddingItsMetric(t *testing.T) {
	//given
	element := NewElement("elementId", "Element Name", "SERVER", "use1a")
	sample, _ := NewSample("metricId", time.Now(), 0.0)
	//when
	element.AddSample(sample)
	//then
	if len(element.samples) != 1 || len(element.Samples()) != 1 {
		t.Errorf("element should contain exactly 1 sample")
	}

	if len(element.metrics) != 1 {
		t.Errorf("element should contain exactly 1 metric")
	}
}

func TestAddElementSampleWontOverrideExistingMetric(t *testing.T) {
	//given
	element := NewElement("elementId", "Element Name", "SERVER", "use1a")
	metric := NewMetric("metricId", "Metric Name", "bytes", "COUNTER", "NONE")
	sample, _ := NewSample("metricId", time.Now(), 0.0)
	//when
	element.AddMetric(metric)
	element.AddSample(sample)
	//then
	if len(element.metrics) != 1 {
		t.Errorf("element should contain exactly 1 metric")
	}
	if m := element.metrics["metricId"]; m.Name != "Metric Name" || m.Unit != "bytes" || m.Type != "COUNTER" || m.SparseDataStrategy != "NONE" {
		t.Errorf("element's existing metric should not be overriden by sample")
	}
}

func TestElementMarshalJSON(t *testing.T) {
	//given
	element := NewElement("elementId", "Element Name", "SERVER", "use1a")
	metric := NewMetric("metricId", "Metric Name", "bytes", "COUNTER", "NONE")
	sample, _ := NewSample("metricId", time.Now(), 0.0)
	element.AddMetric(metric)
	element.AddSample(sample)
	//when
	//marshal
	ejson, _ := element.MarshalJSON()
	log.Println("marshaled element json: ", string(ejson))
	//unmarshal
	var e Element
	e.UnmarshalJSON(ejson)
	if e.Id != "elementId" || e.Name != "Element Name" || e.Type != "SERVER" || e.Location != "use1a" {
		t.Errorf("unmarshaled marshaled element is not equal to its original")
	}

	if len(e.samples) != 1 {
		t.Errorf("element should contain exactly 1 sample")
	}

}