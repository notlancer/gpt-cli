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

type ConversationItemEvent struct {
	Type string           `json:"type"`
	Item conversationItem `json:"item"`
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
