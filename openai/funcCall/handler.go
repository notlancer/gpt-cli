package funcCall

import (
	"fmt"
	"github.com/notlancer/gpt-cli/openai/events"
	funcCallEvents "github.com/notlancer/gpt-cli/openai/funcCall/events"
)

type HandlerStruct struct {
	Callback func(args map[string]any, callID string) (bool, events.ConversationItemEvent)
}

var funcCallHandlers = map[string]HandlerStruct{
	"multiplies_two_numbers": {
		Callback: multipleFunCall,
	},
}

func GetUpdateSessionFunCall() funcCallEvents.SessionUpdateEvent {
	return funcCallEvents.BuildSessionUpdateFuncCallMsg()
}

func Handler(funcCallName string, args map[string]any, callID string) (bool, events.ConversationItemEvent) {
	if handler, ok := funcCallHandlers[funcCallName]; ok {
		return handler.Callback(args, callID)
	}

	return false, events.ConversationItemEvent{}
}

func multipleFunCall(args map[string]any, callID string) (bool, events.ConversationItemEvent) {
	sum := args[funcCallEvents.ArgumentMultipleKeyFirst].(float64) * args[funcCallEvents.ArgumentMultipleKeySecond].(float64)

	return true, events.BuildConvCreateCallFuncMsg(callID, fmt.Sprintf("%f", sum))
}
