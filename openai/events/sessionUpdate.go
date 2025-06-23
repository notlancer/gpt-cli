package events

type tool struct {
	Type        string                 `json:"_type"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
}

type session struct {
	Tools []tool `json:"tools"`
}

type SessionUpdateEvent struct {
	Type    string  `json:"type"`
	Session session `json:"session"`
}

func BuildSessionUpdateMsg() SessionUpdateEvent {
	return SessionUpdateEvent{
		Type: "session.update",
		Session: session{Tools: []tool{
			{Type: "string", Name: "<UNK>", Description: "<UNK>", Parameters: map[string]interface{}{}},
		}},
	}
}
