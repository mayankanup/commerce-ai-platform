package loader_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/rag/loader"
)

func TestLoadDocuments(
	t *testing.T,
) {

	dir := t.TempDir()

	err := os.WriteFile(
		filepath.Join(dir, "faq.md"),
		[]byte("# FAQ\nHello World"),
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(
		filepath.Join(dir, "returns.md"),
		[]byte("# Returns"),
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(
		filepath.Join(dir, "ignore.pdf"),
		[]byte("pdf"),
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}

	loader := loader.NewFilesystemLoader(dir)

	documents, err := loader.Load()
	if err != nil {
		t.Fatal(err)
	}

	if len(documents) != 2 {
		t.Fatalf(
			"expected 2 documents, got %d",
			len(documents),
		)
	}

	if documents[0].Content == "" {
		t.Fatal("document content should not be empty")
	}
}
