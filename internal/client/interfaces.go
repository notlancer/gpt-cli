package client

type OpenAIClient interface {
	SendChatMsg(msg string) error
	StartUserGPTChat() error
	Close() error
}
