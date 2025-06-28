package container

import (
	"github.com/notlancer/gpt-cli/internal/client"
	"github.com/notlancer/gpt-cli/internal/config"
)

type Container struct {
	config       *config.Env
	openAIClient client.OpenAIClient
}

func NewContainer() *Container {
	config := config.NewEnv()
	openAIClient := client.Login(config.GetBearerToken())

	return &Container{
		config:       config,
		openAIClient: openAIClient,
	}
}

func (c *Container) Config() *config.Env {
	return c.config
}

func (c *Container) OpenAIClient() client.OpenAIClient {
	return c.openAIClient
}

func (c *Container) Close() error {
	if c.openAIClient != nil {
		return c.openAIClient.Close()
	}

	return nil
}
