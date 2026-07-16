package loader

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mayankanup/commerce-ai-platform/internal/rag/domain"
)

type FilesystemLoader struct {
	root string
}

func NewFilesystemLoader(
	root string,
) *FilesystemLoader {

	return &FilesystemLoader{
		root: root,
	}
}

func (l *FilesystemLoader) Load() ([]domain.Document, error) {

	var documents []domain.Document

	err := filepath.WalkDir(
		l.root,
		func(
			path string,
			entry fs.DirEntry,
			err error,
		) error {

			if err != nil {
				return err
			}

			if entry.IsDir() {
				return nil
			}

			if !isSupported(path) {
				return nil
			}

			bytes, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf(
					"read %s: %w",
					path,
					err,
				)
			}

			documents = append(
				documents,
				domain.Document{
					Name:    entry.Name(),
					Path:    path,
					Content: string(bytes),
				},
			)

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return documents, nil
}

func isSupported(
	path string,
) bool {

	ext := strings.ToLower(
		filepath.Ext(path),
	)

	switch ext {

	case ".md":
		return true

	case ".txt":
		return true

	default:
		return false
	}
}
