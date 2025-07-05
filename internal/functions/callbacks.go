package functions

import (
	"fmt"
)

func multiplyCallback(params map[string]interface{}) (interface{}, error) {
	var numbers []float64
	for _, value := range params {
		if num, ok := value.(float64); ok {
			numbers = append(numbers, num)
		}
	}

	if len(numbers) != 2 {
		return nil, fmt.Errorf("expected 2 numbers, got %d", len(numbers))
	}

	result := numbers[0] * numbers[1]
	return result, nil
}
