package functions

import (
	"github.com/notlancer/gpt-cli/internal/builders"
)

type HandlerStruct struct {
	Callback func(args map[string]any, handler HandlerStruct, callID string) (bool, builders.ConversationItemEvent)
	Tool     tool
}

type GenericCallback func(params map[string]interface{}) (interface{}, error)

type tool struct {
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  parameters `json:"parameters"`
}

type parameters struct {
	Type       string              `json:"type"`
	Properties map[string]property `json:"properties"`
	Required   []string            `json:"required"`
}

type property struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type GenericHandler struct {
	Callback GenericCallback
	Tool     tool
}
