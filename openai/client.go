package openai

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/notlancer/gpt-cli/openai/events"
	"github.com/notlancer/gpt-cli/openai/funcCall"
	"os"
	"strings"
)

type Client struct {
	Token string
	ws    *websocket.Conn
}

func Login(token string) *Client {
	client := Client{Token: token}
	client.ws = CreateWSConnection(client.Token)

	sessionUpdateEvent := funcCall.GetUpdateSessionFunCall()
	SendWsMessage(client.ws, sessionUpdateEvent)

	ListenToWsMessage(client.ws, &client)

	return &client
}

func (c *Client) SendChatMsg(msg string) {
	startMessage := events.BuildConversationCreateMsg(msg)
	SendWsMessage(c.ws, startMessage)

	responseCreate := events.BuildResponseCreateMsg()
	SendWsMessage(c.ws, responseCreate)
}

func (c *Client) StartUserGPTChat() {
	go func() {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("...: ")

		input, _ := reader.ReadString('\n')
		strings.TrimSpace(input)

		c.SendChatMsg(input)
	}()
}
