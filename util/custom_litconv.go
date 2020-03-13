package util

import (
	"fmt"
	"strconv"
)

/*
FloatValue Convert the literal value of a scanned token to int64
*/
func FloatValue(lit []byte) (float64, error) {
	return strconv.ParseFloat(string(lit), 64)
}

/*
BoolValue Convert the literal value of a scanned token to int64
// @TODO: is this funxction really needed?
*/
func BoolValue(lit []byte) (bool, error) {
	s := string(lit)
	if s == "true" {
		return true, nil
	}
	if s == "false" {
		return false, nil
	}
	return false, fmt.Errorf("%s does not represent a boolean literal", s)
}
