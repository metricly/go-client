package core

import (
	"log"
	"testing"
)

func TestSample(t *testing.T) {
	test := struct {
		metricId string
		timestamp int64
		value float64
	} {
		metricId: "metricId",
		timestamp: 1517873901000,
		value: 0.0,
	}
	sample, _ := NewSample(test.metricId, test.timestamp, test.value)
	log.Println(sample)
	if sample.MetricId() != test.metricId {
		t.Errorf("sample metricId expected: %s, actual:%s", test.metricId, sample.MetricId())
	}
	if int64(sample.Timestamp()) != test.timestamp {
		t.Errorf("sample timestamp expected: %d, actual: %d", test.timestamp, sample.Timestamp())
	}
	if sample.Val() != 0.0 {
		t.Errorf("sample val expected: %f, actual: %f", test.value, sample.Val())
	}
}

func TestSamples(t *testing.T) {
	tests := map[string] struct {
		metricId string
		timestamp interface{}
		value interface{}
		valid bool
	} {
		"string sample with epoch-milliseconds timestamp": {
			metricId: "metricId",
			timestamp: 1517873901000,
			value: "0",
			valid: true,
		},
		"int sample with RFC3339 timestamp": {
			metricId: "metricId",
			timestamp: "1970-01-01T00:00:00Z",
			value: 0,
			valid: true,
		},
	}

	for name, test := range tests {
		sample, err := NewSample(test.metricId, test.timestamp, test.value)
		if err == nil && test.valid == true {
			log.Println("positive test \"" + name + "\" OK", sample)
		} else if err != nil && test.valid == false {
			log.Println("negative test \"" + name + "\" OK", sample, err)
		} else {
			t.Errorf("test %s FAIL, expected: %t, sample: %v, error: %v", name, test.valid, sample, err)
		}
	}
}