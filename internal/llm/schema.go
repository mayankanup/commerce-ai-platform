package llm

// JSONSchema models the subset of JSON Schema
// needed for LLM tool definitions.
type JSONSchema struct {
	Type string

	Properties map[string]Property

	Required []string
}

type Property struct {
	Type string

	Description string

	Enum []string
}
