package llm

// ToolDefinition describes a callable tool exposed to an LLM.
//
// ParametersSchema contains the provider-independent JSON schema
// describing the tool arguments.
//
// The Ollama/OpenAI adapters are responsible for converting this
// structure into their native API representation.
type ToolDefinition struct {
	Name string

	Description string

	ParametersSchema JSONSchema
}

// ToolCall represents one tool invocation requested by the model.
type ToolCall struct {
	Name string

	Arguments map[string]any
}
