package adapter

import "reflect"

// Adapt Adapts CRUD object
func Adapt(o interface{}) map[string]interface{} {
	if reflect.TypeOf(o).Kind() == reflect.Slice {
		aux := reflect.ValueOf(o)
		return map[string]interface{}{
			"total":   aux.Len(),
			"results": o,
		}
	}

	return map[string]interface{}{
		"total":   1,
		"results": []interface{}{o},
	}
}
