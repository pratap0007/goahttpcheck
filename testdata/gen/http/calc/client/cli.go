// Code generated by goa v3.0.9, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen calc/design

package client

import (
	calc "calc/gen/calc"
	"fmt"
	"strconv"
)

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddA string, calcAddB string) (*calc.AddPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	payload := &calc.AddPayload{
		A: a,
		B: b,
	}
	return payload, nil
}

// BuildDivPayload builds the payload for the calc div endpoint from CLI flags.
func BuildDivPayload(calcDivA string, calcDivB string) (*calc.DivPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	payload := &calc.DivPayload{
		A: a,
		B: b,
	}
	return payload, nil
}
