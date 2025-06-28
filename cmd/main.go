package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/notlancer/gpt-cli/internal/container"
)

func main() {
	container := container.NewContainer()
	defer container.Close()

	log.Println("Connected successfully to OpenAI WebSocket")
	fmt.Print("Hello there, I'm ChatGPT, what about you?\n")

	if err := container.OpenAIClient().StartUserGPTChat(); err != nil {
		log.Fatalf("Failed to start chat: %v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// keep the application running
	go func() {
		for {
			time.Sleep(2 * time.Second)
		}
	}()

	<-sigs
	log.Println("Shutting down gracefully...")
}
