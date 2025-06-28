package functions

import (
	"fmt"

	"github.com/notlancer/gpt-cli/internal/builders"
)

const (
	multipleFuncFirstParam  = "number1"
	multipleFuncSecondParam = "number2"
)

type HandlerStruct struct {
	Callback func(args map[string]any, callID string) (bool, builders.ConversationItemEvent)
}

var funcCallHandlers = map[string]HandlerStruct{
	"multiplies_two_numbers": {
		Callback: multipleFunCall,
	},
}

func GetUpdateSessionFunCall() map[string]interface{} {
	return map[string]interface{}{
		"type": "session.update",
		"session": map[string]interface{}{
			"tools": []interface{}{
				map[string]interface{}{
					"type":        "function",
					"name":        "multiplies_two_numbers",
					"description": "multiplies two numbers",
					"parameters": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							multipleFuncFirstParam: map[string]interface{}{
								"type":        "number",
								"description": "The first number.",
							},
							multipleFuncSecondParam: map[string]interface{}{
								"type":        "number",
								"description": "The second number.",
							},
						},
						"required": []interface{}{
							multipleFuncFirstParam,
							multipleFuncSecondParam,
						},
					},
				},
			},
			"tool_choice": "auto",
		},
	}
}

func Handler(funcCallName string, args map[string]any, callID string) (bool, builders.ConversationItemEvent) {
	if handler, ok := funcCallHandlers[funcCallName]; ok {
		return handler.Callback(args, callID)
	}

	return false, builders.ConversationItemEvent{}
}

func multipleFunCall(args map[string]any, callID string) (bool, builders.ConversationItemEvent) {
	first, ok1 := args[multipleFuncFirstParam].(float64)
	second, ok2 := args[multipleFuncSecondParam].(float64)

	if !ok1 || !ok2 {
		return false, builders.BuildConvCreateCallFuncMsg(callID, "Error while parsing multiplies function")
	}

	result := first * second
	return true, builders.BuildConvCreateCallFuncMsg(callID, fmt.Sprintf("%f", result))
}
