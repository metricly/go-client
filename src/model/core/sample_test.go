package core

import (
	"testing"
)

func TestSample(t *testing.T) {
	//given
	test := struct {
		metricId string
		timestamp int64
		value float64
	} {
		metricId: "metricId",
		timestamp: 1517873901000,
		value: 0.0,
	}
	//when
	sample, _ := NewSample(test.metricId, test.timestamp, test.value)
	//then
	if sample.MetricId() != test.metricId {
		t.Errorf("sample metricId expected: %s, actual: %s", test.metricId, sample.MetricId())
	}
	if int64(sample.Timestamp()) != test.timestamp {
		t.Errorf("sample timestamp expected: %d, actual: %d", test.timestamp, sample.Timestamp())
	}
	if sample.Val() != 0.0 {
		t.Errorf("sample val expected: %f, actual: %f", test.value, sample.Val())
	}
}

func TestSampleWithBlankMetricId(t *testing.T) {
	//given & when
	sample, err := NewSample("  ", 1517873901000, 0.0)
	//then
	if err == nil {
		t.Errorf("blank metric id should fail on sample %s", sample)
	}
}

func TestSamplesTypeConversions(t *testing.T) {
	//given
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
	//when & then
	for name, test := range tests {
		sample, err := NewSample(test.metricId, test.timestamp, test.value)
		if err == nil && test.valid == false ||  err != nil && test.valid == true {
			t.Errorf("test %s FAIL, expected: %t, sample: %v, error: %v", name, test.valid, sample, err)
		}
	}
}