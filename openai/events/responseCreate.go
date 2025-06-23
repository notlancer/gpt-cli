package events

type response struct {
	Modalities []string `json:"modalities"`
}

type ResponseCreateEvent struct {
	Type     string   `json:"type"`
	Response response `json:"response"`
}

func BuildResponseCreateMsg() ResponseCreateEvent {
	return ResponseCreateEvent{
		Type: "response.create",
		Response: response{
			Modalities: []string{"text"},
		},
	}
}
