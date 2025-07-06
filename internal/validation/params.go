package validation

import (
	"fmt"
	"reflect"
)

type ValidatedParams struct {
	params map[string]interface{}
}

func (vp *ValidatedParams) GetString(key string) string {
	return vp.params[key].(string)
}

func ValidateRequiredParams(message map[string]interface{}, required map[string]reflect.Type) (*ValidatedParams, error) {
	validated := &ValidatedParams{
		params: make(map[string]interface{}),
	}

	for paramName, expectedType := range required {
		value, exists := message[paramName]
		if !exists {
			return nil, fmt.Errorf("required parameter '%s' is missing", paramName)
		}

		if value == nil {
			return nil, fmt.Errorf("required parameter '%s' is null", paramName)
		}

		actualType := reflect.TypeOf(value)
		if actualType != expectedType {
			return nil, fmt.Errorf("parameter '%s' has wrong type: expected %v, got %v",
				paramName, expectedType, actualType)
		}

		validated.params[paramName] = value
	}

	return validated, nil
}
