package main

import (
	"fmt"
	"github.com/notlancer/gpt-cli/bootstrap"
	"github.com/notlancer/gpt-cli/openai"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := bootstrap.CreateApp()
	client := openai.Login(app.Env.BearerToken)

	log.Println("connected successfully to openai ws")
	fmt.Print("Hello there, im ChatGPT, wbu?\n")

	go openai.RequestUserInput(client)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			time.Sleep(2 * time.Second)
		}
	}()

	<-sigs
}
