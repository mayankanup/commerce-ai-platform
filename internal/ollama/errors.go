package ollama

import "fmt"

type ErrorResponse struct {
	Message string `json:"error"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("ollama: %s", e.Message)
}
