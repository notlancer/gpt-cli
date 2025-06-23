package events

type SessionUpdateEvent struct {
	Type    string  `json:"type"`
	Session Session `json:"session"`
}

type Session struct {
	Tools      []Tool `json:"tools"`
	ToolChoice string `json:"tool_choice"`
}

type Tool struct {
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

type Parameters struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
}

type Property struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

func BuildSessionUpdateMsg() SessionUpdateEvent {
	return SessionUpdateEvent{
		Type: "session.update",
		Session: Session{
			ToolChoice: "auto",
			Tools: []Tool{
				{
					Type:        "function",
					Name:        "multiplies_two_numbers",
					Description: "multiplies two numbers",
					Parameters: Parameters{
						Type: "object",
						Properties: map[string]Property{
							"number1": {
								Type:        "number",
								Description: "The first number.",
							},
							"number2": {
								Type:        "number",
								Description: "The second number.",
							},
						},
						Required: []string{"number1", "number2"},
					},
				},
			},
		},
	}

}
