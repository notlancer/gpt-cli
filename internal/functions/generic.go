package functions

import (
	"fmt"

	"github.com/notlancer/gpt-cli/internal/builders"
)

func createGenericHandler(callback GenericCallback) CallbackFunc {
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

		return true, builders.BuildConvCreateCallFuncMsg(callID, result)
	}
}
