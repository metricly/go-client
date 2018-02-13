package core

import (
	"time"
	"fmt"
	"math"
)

//Timestamp is number of milliseconds since epoch
type Timestamp int64

//TimestampValue converts/casts a supported value type to Timestamp or reports an Error, supported types are:
//	{int64, int, time.RFC3339 string}
func TimestampValue (timestamp interface{}) (Timestamp, error) {
	epoch, err := parseTimestampValue(timestamp)
	return Timestamp(epoch), err
}

func parseTimestampValue(timestamp interface{}) (epoch int64, err error) {
	switch timestamp := timestamp.(type) {
	case int64:
		epoch = timestamp
	case int:
		epoch = int64(timestamp)
	case string:
		{
			ts, err := time.Parse(time.RFC3339, timestamp)
			if err != nil {
				return math.MinInt64, fmt.Errorf("%v is not a supported timestamp value", timestamp)
			}
			epoch = ts.UnixNano() / int64(time.Millisecond)
		}
	case time.Time:
		epoch = timestamp.UnixNano() / int64(time.Millisecond)
	default:
		return math.MinInt64, fmt.Errorf("%v is not a supported timestamp value", timestamp)
	}
	return epoch, nil
}