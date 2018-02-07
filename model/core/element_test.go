package core

import (
	"testing"
	"time"
)

func TestAddElementSampleAddingItsMetric(t *testing.T) {
	//given
	element := NewElement("elementId")
	sample, _ := NewSample("metricId", time.Now(), 0.0)
	//when
	element.AddSample(sample)
	//then
	if len(element.samples) != 1 {
		t.Errorf("element should contain exactly 1 sample")
	}

	if len(element.metrics) != 1 {
		t.Errorf("element should contain exactly 1 metric")
	}
}

func TestAddElementSampleWontOverrideExistingMetric(t *testing.T) {
	//given
	element := NewElement("elementId")
	metric := NewMetric("metricId", "Metric Name", "bytes", "SERVER", "NONE")
	sample, _ := NewSample("metricId", time.Now(), 0.0)
	//when
	element.AddMetric(metric)
	element.AddSample(sample)
	//then
	if len(element.metrics) != 1 {
		t.Errorf("element should contain exactly 1 metric")
	}
	if m := element.metrics["metricId"]; m.Name != "Metric Name" || m.Unit != "bytes" || m.Type != "SERVER" || m.SparseDataStrategy != "NONE" {
		t.Errorf("element's existing metric should not be overriden by sample")
	}

}