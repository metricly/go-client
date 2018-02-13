package core

import (
	"fmt"
	"strings"
	"encoding/json"
)

//Sample represents a single data point for a Metric, use its Construction function NewSample to create a Sample
type Sample struct {
	metricId string
	timestamp Timestamp
	val Number
}

type sampleKey struct {
	metricId string
	timestamp Timestamp
}

//NewSample constructs a Sample given its metricId, timestamp and value or reports an Error
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

//returns sample metricId
func (s *Sample) MetricId() string {
	return s.metricId
}

//returns sample timestamp
func (s *Sample) Timestamp() int64 {
	return int64(s.timestamp)
}

//returns sample value
func (s *Sample) Val() float64 {
	return float64(s.val)
}

func (s *Sample) key() sampleKey {
	return sampleKey{s.metricId, s.timestamp}
}

func (s Sample) String() string {
	return fmt.Sprintf("%s => %.3f @[%d]", s.metricId, s.val, s.timestamp)
}

type sampleJSON struct {
	MetricId string `json:"metricId"`
	Timestamp int64 `json:"timestamp"`
	Val float64	`json:"val"`
}

func (s Sample) MarshalJSON() ([]byte, error) {
	sjson := sampleJSON{
		s.MetricId(),
		s.Timestamp(),
		s.Val(),
	}
	return json.Marshal(sjson)
}

func (s *Sample) UnmarshalJSON(b []byte) error {
	var sjson sampleJSON
	if err := json.Unmarshal(b, &sjson); err != nil {
		return err
	}
	*s, _ = NewSample(sjson.MetricId, sjson.Timestamp, sjson.Val)
	return nil
}