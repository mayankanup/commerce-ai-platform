package agent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/llm"
)

type Options struct {
	MaxToolRounds int
}

type Agent struct {
	llm      llm.Client
	registry *Registry

	maxToolRounds int
}

func New(
	client llm.Client,
	registry *Registry,
	options Options,
) *Agent {

	maxRounds := options.MaxToolRounds
	if maxRounds <= 0 {
		maxRounds = 5
	}

	return &Agent{
		llm:           client,
		registry:      registry,
		maxToolRounds: maxRounds,
	}
}

// Chat is the public entry point into the agent.
func (a *Agent) Chat(
	ctx context.Context,
	prompt string,
) (*ChatResult, error) {

	messages := []llm.Message{
		{
			Role:    llm.UserRole,
			Content: prompt,
		},
	}

	return a.run(ctx, messages)
}

// run executes the complete LLM <-> Tool orchestration loop.
func (a *Agent) run(
	ctx context.Context,
	messages []llm.Message,
) (*ChatResult, error) {

	for round := 0; round < a.maxToolRounds; round++ {

		reply, err := a.llm.Chat(
			ctx,
			messages,
			a.registry.Definitions(),
		)

		if err != nil {
			return nil, err
		}

		// Keep the assistant message.
		messages = append(messages, *reply)

		// No tools requested. Conversation is complete.
		if len(reply.ToolCalls) == 0 {

			return &ChatResult{
				Response: reply.Content,
			}, nil
		}

		// Execute every requested tool.
		for _, toolCall := range reply.ToolCalls {

			toolMessage, err := a.executeToolCall(
				ctx,
				toolCall,
			)

			if err != nil {
				return nil, err
			}

			messages = append(messages, toolMessage)
		}
	}

	return nil, fmt.Errorf(
		"maximum tool rounds (%d) exceeded",
		a.maxToolRounds,
	)
}

func (a *Agent) executeToolCall(
	ctx context.Context,
	call llm.ToolCall,
) (llm.Message, error) {

	tool, err := a.registry.Get(call.Name)
	if err != nil {
		return llm.Message{}, err
	}

	result, err := tool.Execute(
		ctx,
		call.Arguments,
	)
	if err != nil {
		return llm.Message{}, err
	}

	payload, err := json.Marshal(result.Content)
	if err != nil {
		return llm.Message{}, err
	}

	return llm.Message{
		Role:     llm.ToolRole,
		ToolName: call.Name,
		Content:  string(payload),
	}, nil
}
