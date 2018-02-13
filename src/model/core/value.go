package core

import (
	"math"
	"fmt"
	"strconv"
)

type Number float64

//NumericValue converts/casts a supported value type to a Number or reports an Error, supported types are:
//	{float64, float32, int64, int32, int16, int8, int, string}
func NumericValue(value interface{}) (Number, error) {
	amount, err := parseNumericValue(value)
	return Number(amount), err
}

func parseNumericValue(value interface{}) (amount float64, err error) {
	switch value := value.(type) {
	case float64:
		amount = value
	case float32:
		amount = float64(value)
	case int64:
		amount = float64(value)
	case int32:
		amount = float64(value)
	case int16:
		amount = float64(value)
	case int8:
		amount = float64(value)
	case int:
		amount = float64(value)
	case string:
		{
			if val, err := strconv.ParseFloat(value, 64); err != nil {
				return math.NaN(), fmt.Errorf("%s is not a valid number string, error: ", value, err)
			} else {
				amount = val
			}
		}

	default:
		return math.NaN(), fmt.Errorf("%v is not a supported number type", value)
	}
	return amount, nil
}