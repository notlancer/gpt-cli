package events

type conversationItemContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type conversationItem struct {
	Type    string                    `json:"type"`
	Role    string                    `json:"role"`
	Content []conversationItemContent `json:"content"`
}

type conversationItemFuncCall struct {
	Type   string `json:"type"`
	CallID string `json:"call_id"`
	Output string `json:"output"`
}

type ConversationItemEvent struct {
	Type string      `json:"type"`
	Item interface{} `json:"item"`
}

func BuildConversationCreateMsg(userInput string) ConversationItemEvent {
	return ConversationItemEvent{
		Type: "conversation.item.create",
		Item: conversationItem{
			Type: "message",
			Role: "user",
			Content: []conversationItemContent{
				{Type: "input_text", Text: userInput},
			},
		},
	}
}

func BuildConvCreateCallFuncMsg(callId string, output string) ConversationItemEvent {
	return ConversationItemEvent{
		Type: "conversation.item.create",
		Item: conversationItemFuncCall{
			Type:   "function_call_output",
			CallID: callId,
			Output: output,
		},
	}
}
