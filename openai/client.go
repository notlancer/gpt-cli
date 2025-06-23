package openai

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/notlancer/gpt-cli/openai/events"
	"log"
	"os"
	"strings"
)

const (
	WsScheme = "wss"
	WsHost   = "api.openai.com"
	WsPath   = "/v1/realtime"
	WsQuery  = "model=gpt-4o-mini-realtime-preview-2024-12-17"

	AuthorizationHeaderKey = "Authorization"
	OpenaiBetaHeaderKey    = "OpenAI-Beta"
	OpenaiBetaHeaderValue  = "realtime=v1"
)

type Client struct {
	Token string
	Ws    *websocket.Conn
}

func Login(token string) *Client {
	client := Client{Token: token}
	client.Ws = client.createWSConnection()

	return &client
}

func RequestUserInput(client *Client) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("...: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	startMessage := events.BuildConversationCreateMsg(input)
	client.SendWsMessage(startMessage)

	responseCreate := events.BuildResponseCreateMsg()
	client.SendWsMessage(responseCreate)
}

func (c *Client) SendWsMessage(msg interface{}) {
	rawMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = c.Ws.WriteMessage(websocket.TextMessage, rawMsg)
	if err != nil {
		log.Fatal(err)
		return
	}
}
