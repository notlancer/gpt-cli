package functions

import (
	"github.com/notlancer/gpt-cli/internal/builders"
)

func GetUpdateSessionFunCall() map[string]interface{} {
	tools := make([]interface{}, 0, len(funcCallHandlers))

	for _, handler := range funcCallHandlers {
		tools = append(tools, handler.Tool)
	}

	return map[string]interface{}{
		"type": "session.update",
		"session": map[string]interface{}{
			"tools":       tools,
			"tool_choice": "auto",
		},
	}
}

func Handler(funcCallName string, args map[string]any, callID string) (bool, builders.ConversationItemEvent) {
	if handler, ok := funcCallHandlers[funcCallName]; ok {
		return handler.Callback(args, handler, callID)
	}

	return false, builders.ConversationItemEvent{}
}
