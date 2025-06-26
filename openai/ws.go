package openai

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

const (
	WsScheme = "wss"
	WsHost   = "api.openai.com"
	WsPath   = "/v1/realtime"
	WsQuery  = "model=gpt-4o-mini-realtime-preview-2024-12-17"

	AuthorizationHeaderKey = "Authorization"
	OpenAIBetaHeaderKey    = "OpenAI-Beta"
	OpenaiBetaHeaderValue  = "realtime=v1"
)

func CreateWSConnection(token string) *websocket.Conn {
	wsUrl := url.URL{Scheme: WsScheme, Host: WsHost, Path: WsPath, RawQuery: WsQuery}
	ws, _, err := websocket.DefaultDialer.Dial(wsUrl.String(), http.Header{
		AuthorizationHeaderKey: {fmt.Sprintf("Bearer %s", token)},
		OpenAIBetaHeaderKey:    {OpenaiBetaHeaderValue},
	})

	if err != nil {
		log.Fatal("dial:", err)
	}

	return ws
}

func SendWsMessage(ws *websocket.Conn, msg interface{}) {
	rawMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, rawMsg)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ListenToWsMessage(ws *websocket.Conn, client *Client) {
	go func() {
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Fatal("read:", err)
				return
			}

			HandleWsMessage(client, message)
		}
	}()
}
