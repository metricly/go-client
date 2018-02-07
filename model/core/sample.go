package core

import (
	"fmt"
	"strings"
	"encoding/json"
)

type Sample struct {
	metricId string
	timestamp Timestamp
	val Number
}

type SampleKey struct {
	metricId string
	timestamp Timestamp
}

func NewSample(metricId string, timestamp interface{}, value interface{}) (Sample, error) {
	var errors []string
	if len(strings.TrimSpace(metricId)) == 0 {
		errors = append(errors, "metricId can't be blank")
	}
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

func (s *Sample) Timestamp() int64 {
	return int64(s.timestamp)
}

func (s *Sample) Val() float64 {
	return float64(s.val)
}

func (s *Sample) Key() SampleKey {
	return SampleKey{s.metricId, s.timestamp}
}

func (s Sample) String() string {
	return fmt.Sprintf("%s => %.3f @[%d]", s.metricId, s.val, s.timestamp)
}

type SampleJSON struct {
	MetricId string
	Timestamp int64
	Val float64
}

func (s Sample) MarshalJSON() ([]byte, error) {
	sjson := SampleJSON {
		s.MetricId(),
		s.Timestamp(),
		s.Val(),
	}
	return json.Marshal(sjson)
}

func (s *Sample) UnmarshalJSON(b []byte) error {
	var sjson SampleJSON
	if err := json.Unmarshal(b, &sjson); err != nil {
		return err
	}
	*s, _ = NewSample(sjson.MetricId, sjson.Timestamp, sjson.Val)
	return nil
}