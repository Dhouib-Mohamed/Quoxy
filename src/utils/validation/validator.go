package validation

import (
	"fmt"
	"strconv"
)

type field struct {
	Name     string
	Type     string
	Optional bool
}

var transformFn = map[string]func(string) any{
	"number": func(s string) any {
		n, _ := strconv.Atoi(s)
		return n
	},
	"float": func(s string) any {
		f, _ := strconv.ParseFloat(s, 64)
		return f
	},
	"string": func(s string) any {
		return s
	},
}

func validation(data map[string]any, fields []field) error {
	for _, field := range fields {
		name := field.Name
		dataType := field.Type
		optional := field.Optional

		dataValue, ok := data[name]
		if !ok {
			if optional {
				continue
			} else {
				return fmt.Errorf("field %s is required", name)
			}
		}

		switch dataValue.(type) {
		case string:
			if dataType != "string" {
				data[name] = transformFn[dataType](dataValue.(string))
			}
		default:
			return fmt.Errorf("field %s must be a string", name)
		}
	}
	return nil
}
