package builders

type ConversationItemContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ConversationItem struct {
	Type    string                    `json:"type"`
	Role    string                    `json:"role"`
	Content []ConversationItemContent `json:"content"`
}

type ConversationItemFuncCall struct {
	Type   string `json:"type"`
	CallID string `json:"call_id"`
	Output string `json:"output"`
}

type ConversationItemEvent struct {
	Type string      `json:"type"`
	Item interface{} `json:"item"`
}

type ResponseCreateEvent struct {
	Type     string `json:"type"`
	Response struct {
		Modalities []string `json:"modalities"`
	} `json:"response"`
}

func BuildConversationCreateMsg(userInput string) ConversationItemEvent {
	return ConversationItemEvent{
		Type: "conversation.item.create",
		Item: ConversationItem{
			Type: "message",
			Role: "user",
			Content: []ConversationItemContent{
				{Type: "input_text", Text: userInput},
			},
		},
	}
}

func BuildConvCreateCallFuncMsg(callId string, output string) ConversationItemEvent {
	return ConversationItemEvent{
		Type: "conversation.item.create",
		Item: ConversationItemFuncCall{
			Type:   "function_call_output",
			CallID: callId,
			Output: output,
		},
	}
}

func BuildResponseCreateMsg() ResponseCreateEvent {
	return ResponseCreateEvent{
		Type: "response.create",
		Response: struct {
			Modalities []string `json:"modalities"`
		}{Modalities: []string{"text"}},
	}
}
