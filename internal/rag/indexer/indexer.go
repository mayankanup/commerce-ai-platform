package indexer

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/google/uuid"

	"github.com/mayankanup/commerce-ai-platform/internal/embedding"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/chunker"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/domain"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/loader"
	"github.com/mayankanup/commerce-ai-platform/internal/rag/repository"
)

type Indexer struct {
	loader *loader.FilesystemLoader

	chunker *chunker.Chunker

	embedding embedding.Provider

	repository repository.Repository
}

func New(
	loader *loader.FilesystemLoader,
	chunker *chunker.Chunker,
	embedding embedding.Provider,
	repository repository.Repository,
) *Indexer {

	return &Indexer{
		loader:     loader,
		chunker:    chunker,
		embedding:  embedding,
		repository: repository,
	}
}

func (i *Indexer) Index(
	ctx context.Context,
) error {

	documents, err := i.loader.Load()
	if err != nil {
		return err
	}

	for _, document := range documents {

		if err := i.indexDocument(
			ctx,
			document,
		); err != nil {
			return err
		}
	}

	return nil
}

func (i *Indexer) Rebuild(
	ctx context.Context,
) error {

	if err := i.repository.DeleteAll(ctx); err != nil {
		return err
	}

	return i.Index(ctx)
}

func (i *Indexer) Count(
	ctx context.Context,
) (int, error) {

	return i.repository.Count(ctx)
}

func (i *Indexer) indexDocument(
	ctx context.Context,
	document domain.Document,
) error {

	// Remove previously indexed chunks for this document.
	logger.Info(
		"Indexing",
		"document", document.Name,
	)
	if err := i.repository.Delete(
		ctx,
		document.Name,
	); err != nil {
		return err
	}

	chunks := i.chunker.Chunk(document)

	vectors := make(
		[]repository.DocumentVector,
		0,
		len(chunks),
	)

	for _, chunk := range chunks {

		vector, err := i.embedding.Embed(
			ctx,
			chunk.Content,
		)
		if err != nil {
			return err
		}

		vectors = append(
			vectors,
			repository.DocumentVector{
				ID: uuid.NewString(),

				DocumentID: document.Name,

				ChunkID: chunk.ID,

				Content: chunk.Content,

				Embedding: vector,

				Metadata: map[string]string{
					"name":  document.Name,
					"path":  document.Path,
					"title": chunk.Title,
				},
			},
		)
	}

	return i.repository.Upsert(
		ctx,
		vectors,
	)
}
