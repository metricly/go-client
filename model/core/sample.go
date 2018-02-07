package core

import (
	"fmt"
	"strings"
)

type Sample struct {
	metricId string
	timestamp Timestamp
	val Number
}

func NewSample(metricId string, timestamp interface{}, value interface{}) (Sample, error) {
	var errors []string
	ts, timestampError := TimestampValue(timestamp)
	if timestampError != nil {
		errors = append(errors, timestampError.Error())
	}
	val, valueError := NumericValue(value)
	if valueError != nil {
		errors = append(errors, valueError.Error())
	}
	if len(errors) > 0 {
		return Sample{metricId, ts, val}, fmt.Errorf(strings.Join(errors, ";"))
	}
	return Sample{metricId, ts, val}, nil
}


func (s *Sample) MetricId() string {
	return s.metricId
}

func (s *Sample) Timestamp() Timestamp {
	return s.timestamp
}

func (s *Sample) Val() float64 {
	return float64(s.val)
}

func (s Sample) String() string {
	return fmt.Sprintf("%s => %.3f @[%d]", s.metricId, s.val, s.timestamp)
}