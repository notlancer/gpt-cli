package interfaces

type MessageSender interface {
	SendMessage(msg interface{}) error
}

type MessageHandler interface {
	HandleMessage(message []byte) error
}

type ChatController interface {
	StartUserGPTChat() error
}

type MessageClient interface {
	MessageSender
	ChatController
}
