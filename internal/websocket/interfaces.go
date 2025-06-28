package websocket

import "github.com/notlancer/gpt-cli/internal/interfaces"

type WebSocketClient interface {
	interfaces.MessageSender
	ListenForMessages(client interfaces.MessageHandler) error
	Close() error
}
