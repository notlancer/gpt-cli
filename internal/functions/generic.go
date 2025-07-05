package functions

import (
	"fmt"

	"github.com/notlancer/gpt-cli/internal/builders"
)

func createGenericHandler(callback GenericCallback) func(args map[string]any, handler HandlerStruct, callID string) (bool, builders.ConversationItemEvent) {
	return func(args map[string]any, handler HandlerStruct, callID string) (bool, builders.ConversationItemEvent) {
		extractedParams := make(map[string]interface{})
		for paramName := range handler.Tool.Parameters.Properties {
			if value, exists := args[paramName]; exists {
				extractedParams[paramName] = value
			}
		}

		result, err := callback(extractedParams)
		if err != nil {
			return false, builders.BuildConvCreateCallFuncMsg(callID, fmt.Sprintf("Error: %v", err))
		}

		var resultStr string
		switch v := result.(type) {
		case string:
			resultStr = v
		case float64:
			resultStr = fmt.Sprintf("%f", v)
		case int:
			resultStr = fmt.Sprintf("%d", v)
		default:
			resultStr = fmt.Sprintf("%v", v)
		}

		return true, builders.BuildConvCreateCallFuncMsg(callID, resultStr)
	}
}
