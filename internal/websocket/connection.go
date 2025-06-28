package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/notlancer/gpt-cli/internal/interfaces"
)

const (
	WsScheme = "wss"
	WsHost   = "api.openai.com"
	WsPath   = "/v1/realtime"
	WsQuery  = "model=gpt-4o-mini-realtime-preview-2024-12-17"

	AuthorizationHeaderKey = "Authorization"
	OpenAIBetaHeaderKey    = "OpenAI-Beta"
	OpenaiBetaHeaderValue  = "realtime=v1"

	MaxRetries = 3
	RetryDelay = 2 * time.Second
)

type Connection struct {
	conn *websocket.Conn
}

func NewConnection(token string) (*Connection, error) {
	for attempt := 1; attempt <= MaxRetries; attempt++ {
		conn, err := createConnection(token)
		if err == nil {
			return &Connection{conn: conn}, nil
		}

		log.Printf("Connection attempt %d failed: %v", attempt, err)
		if attempt < MaxRetries {
			time.Sleep(RetryDelay)
		}
	}

	return nil, fmt.Errorf("failed after %d attempts", MaxRetries)
}

func createConnection(token string) (*websocket.Conn, error) {
	wsUrl := url.URL{Scheme: WsScheme, Host: WsHost, Path: WsPath, RawQuery: WsQuery}

	headers := http.Header{
		AuthorizationHeaderKey: {fmt.Sprintf("Bearer %s", token)},
		OpenAIBetaHeaderKey:    {OpenaiBetaHeaderValue},
	}

	conn, _, err := websocket.DefaultDialer.Dial(wsUrl.String(), headers)
	return conn, err
}

func (c *Connection) SendMessage(msg any) error {
	rawMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	if err2 := c.conn.WriteMessage(websocket.TextMessage, rawMsg); err2 != nil {
		return fmt.Errorf("failed to send message: %w", err2)
	}

	return nil
}

func (c *Connection) ListenForMessages(client interfaces.MessageHandler) error {
	go func() {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading message: %v", err)
				return
			}

			if err := client.HandleMessage(message); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}()

	return nil
}

func (c *Connection) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}
