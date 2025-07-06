package functions

import (
	"fmt"
)

func multiplyCallback(params map[string]interface{}) (string, error) {
	var numbers []float64
	for _, value := range params {
		if num, ok := value.(float64); ok {
			numbers = append(numbers, num)
		}
	}

	if len(numbers) != 2 {
		return "", fmt.Errorf("expected 2 numbers, got %d", len(numbers))
	}

	result := numbers[0] * numbers[1]
	return fmt.Sprintf("%f", result), nil
}
