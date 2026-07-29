package eval

import (
	"context"
	"strings"
	"testing"
)

func TestCompanyCEO001(t *testing.T) {

	result, err := testApp.Agent.Chat(
		context.Background(),
		"Who is the CEO?",
	)

	if err != nil {
		t.Fatalf("chat failed: %v", err)
	}

	answer := strings.ToLower(result.Response)

	expected := []string{
		"anup mayank",
		"ceo",
	}

	for _, value := range expected {
		if !strings.Contains(answer, value) {
			t.Fatalf(
				"expected answer to contain %q\n\nActual:\n%s",
				value,
				result.Response,
			)
		}
	}
}
