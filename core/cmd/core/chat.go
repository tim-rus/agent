package main

import (
	"github.com/openai/openai-go/v3"
)

//

type Role string

const (
	RoleSystem    Role = "system"
	RoleAssistant Role = "assistant"
	RoleUser      Role = "user"
)

type Message struct {
	Role    Role
	Message string
}

type Usage struct {
	Total     uint64
	Promts    uint64
	Responses uint64
}

//

type Chat struct {
	client  openai.Client
	history []Message
	usage   Usage
	model   string
}

func NewChat(client openai.Client, model string, systemPromt string) *Chat {
	return &Chat{
		client: client,
		history: []Message{
			{
				Role:    RoleSystem,
				Message: systemPromt,
			},
		},
		usage: Usage{},
		model: model,
	}
}
