package openai

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

func (c *Client) createWSConnection() *websocket.Conn {
	wsUrl := url.URL{Scheme: WsScheme, Host: WsHost, Path: WsPath, RawQuery: WsQuery}
	ws, _, err := websocket.DefaultDialer.Dial(wsUrl.String(), http.Header{
		AuthorizationHeaderKey: {fmt.Sprintf("Bearer %s", c.Token)},
		OpenaiBetaHeaderKey:    {OpenaiBetaHeaderValue},
	})

	if err != nil {
		log.Fatal("dial:", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Fatal("read:", err)
				return
			}

			HandleWsMessage(c, message)
		}
	}()

	return ws
}
