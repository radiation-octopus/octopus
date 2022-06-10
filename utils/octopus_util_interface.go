package utils

import (
	"reflect"
	"strconv"
	"strings"
)

func AssignInterface(assign interface{}, value string) interface{} {
	vType := reflect.TypeOf(&assign).String()
	switch vType {
	case "*string":
		return value
	case "*int":
		i, _ := strconv.Atoi(value)
		return i
	case "*float64":
		ss := strings.Split(value, ".")
		f, _ := strconv.ParseFloat(value, len(ss[1]))
		return f
	case "*bool":
		b, _ := strconv.ParseBool(value)
		return b
	default:
		return value
	}

}
