package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/notlancer/gpt-cli/internal/builders"
	"github.com/notlancer/gpt-cli/internal/functions"
	"github.com/notlancer/gpt-cli/internal/interfaces"
	"github.com/notlancer/gpt-cli/internal/messages"
	"github.com/notlancer/gpt-cli/internal/websocket"
)

type Client struct {
	Token     string
	ws        websocket.WebSocketClient
	processor *messages.MessageProcessor
}

var _ OpenAIClient = (*Client)(nil)
var _ interfaces.MessageHandler = (*Client)(nil)
var _ interfaces.MessageClient = (*Client)(nil)

func Login(token string) *Client {
	client := &Client{Token: token}

	wsConn, err := websocket.NewConnection(token)
	if err != nil {
		log.Fatalf("Failed to create WebSocket connection: %v", err)
	}
	client.ws = wsConn

	client.processor = messages.NewMessageProcessor(client)

	sessionUpdateEvent := functions.GetUpdateSessionFunCall()
	if err := client.ws.SendMessage(sessionUpdateEvent); err != nil {
		log.Printf("Warning: Failed to send session update: %v", err)
	}

	if err := client.ws.ListenForMessages(client); err != nil {
		log.Fatalf("Failed to start message listener: %v", err)
	}

	return client
}

func (c *Client) SendChatMsg(msg string) error {
	startMessage := builders.BuildConversationCreateMsg(msg)
	if err := c.ws.SendMessage(startMessage); err != nil {
		return fmt.Errorf("fail: %w", err)
	}

	responseCreate := builders.BuildResponseCreateMsg()
	if err := c.ws.SendMessage(responseCreate); err != nil {
		return fmt.Errorf("fail: %w", err)
	}

	return nil
}

func (c *Client) StartUserGPTChat() error {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("...: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err := c.SendChatMsg(input); err != nil {
			log.Fatalf("Error sending chat message: %v", err)
		}
	}()

	return nil
}

func (c *Client) Close() error {
	if c.ws != nil {
		return c.ws.Close()
	}
	return nil
}

func (c *Client) HandleMessage(message []byte) error {
	return c.processor.ProcessMessage(message)
}

func (c *Client) SendMessage(msg interface{}) error {
	return c.ws.SendMessage(msg)
}
