package chunker

import (
	"fmt"
	"strings"

	"github.com/mayankanup/commerce-ai-platform/internal/rag/domain"
)

const (
	DefaultChunkSize    = 500
	DefaultChunkOverlap = 100
)

type Chunker struct {
	chunkSize int
	overlap   int
}

func New() *Chunker {

	return &Chunker{
		chunkSize: DefaultChunkSize,
		overlap:   DefaultChunkOverlap,
	}
}

func NewWithOptions(
	chunkSize int,
	overlap int,
) *Chunker {

	return &Chunker{
		chunkSize: chunkSize,
		overlap:   overlap,
	}
}

func (c *Chunker) Chunk(
	document domain.Document,
) []domain.DocumentChunk {

	text := normalize(document.Content)

	if text == "" {
		return nil
	}

	var chunks []domain.DocumentChunk

	start := 0
	sequence := 0

	for start < len(text) {

		end := start + c.chunkSize

		if end >= len(text) {
			end = len(text)
		} else {

			for end > start && text[end] != ' ' {
				end--
			}

			if end == start {
				end = start + c.chunkSize
			}
		}

		content := strings.TrimSpace(text[start:end])

		chunks = append(
			chunks,
			domain.DocumentChunk{
				ID: fmt.Sprintf(
					"%s-%d",
					document.Name,
					sequence,
				),
				Document: document.Name,
				Title:    document.Name,
				Content:  content,
				Sequence: sequence,
			},
		)

		if end == len(text) {
			break
		}

		start = end - c.overlap

		if start < 0 {
			start = 0
		}

		sequence++
	}

	return chunks
}

func normalize(
	text string,
) string {

	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\t", " ")

	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}

	for strings.Contains(text, "\n\n\n") {
		text = strings.ReplaceAll(text, "\n\n\n", "\n\n")
	}

	return strings.TrimSpace(text)
}
