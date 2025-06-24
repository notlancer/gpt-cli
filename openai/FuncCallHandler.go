package openai

import (
	"github.com/notlancer/gpt-cli/openai/events"
	"strconv"
)

type FuncCallHandlerStruct struct {
	Callback func(client *Client, args map[string]any, callID string)
}

var funcCallHandlers = map[string]FuncCallHandlerStruct{
	"multiplies_two_numbers": {
		Callback: multipleFunCall,
	},
}

func FuncCallHandler(client *Client, funcCallName string, args map[string]any, callID string) {
	if handler, ok := funcCallHandlers[funcCallName]; ok {
		handler.Callback(client, args, callID)
	}
}

func multipleFunCall(client *Client, args map[string]any, callID string) {
	sum := args["number1"].(float64) * args["number2"].(float64)

	responseMsg := events.BuildConvCreateCallFuncMsg(callID, strconv.FormatFloat(sum, 'f', -1, 64))

	client.SendWsMessage(responseMsg)

	responseCreate := events.BuildResponseCreateMsg()
	client.SendWsMessage(responseCreate)
}
