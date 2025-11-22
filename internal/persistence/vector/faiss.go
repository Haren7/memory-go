package vector

import (
	"context"
	"errors"
	"fmt"
	"memory/internal/embedding"
	"os"
	"path/filepath"

	"github.com/DataIntelligenceCrew/go-faiss"
)

var ErrFaissIndexDoesNotExist = errors.New("faiss index does not exist")

type FaissSearchResponse struct {
	Distances []float32
	Labels    []int64
}

type FaissClient struct {
	dir string
}

func NewFaissClient() *FaissClient {
	return &FaissClient{
		dir: "./faiss-index",
	}
}

func (r *FaissClient) Index(ctx context.Context, indexName string, id int, embedding embedding.Embedding) error {
	if r.exists(indexName) {
		// Load existing index and add the new vector
		index, err := r.load(indexName)
		if err != nil {
			return fmt.Errorf("error loading existing index: %w", err)
		}

		err = index.AddWithIDs(embedding.Vector, []int64{int64(id)})
		if err != nil {
			return fmt.Errorf("error adding vector to existing index: %w", err)
		}

		// Save the updated index
		return r.write(index, indexName)
	}
	err := os.MkdirAll(r.dir, 0755)
	if err != nil {
		return fmt.Errorf("faiss: error creating directory %s: %w", r.dir, err)
	}
	dimension := embedding.Dim
	index, err := faiss.IndexFactory(dimension, "IDMap,Flat", 1)
	if err != nil {
		return fmt.Errorf("error creating idmap + flat index with dim %d - %w", dimension, err)
	}
	err = index.AddWithIDs(embedding.Vector, []int64{int64(id)})
	if err != nil {
		return fmt.Errorf("error adding vector to index: %d - %w", id, err)
	}
	return r.write(index, indexName)
}

func (r *FaissClient) Search(ctx context.Context, indexName string, query embedding.Embedding, topK int) (FaissSearchResponse, error) {
	exists := r.exists(indexName)
	if !exists {
		return FaissSearchResponse{}, ErrFaissIndexDoesNotExist
	}

	index, err := r.load(indexName)
	if err != nil {
		return FaissSearchResponse{}, fmt.Errorf("faiss: error loading index: %w", err)
	}

	distances, labels, err := index.Search(query.Vector, int64(topK))
	if err != nil {
		return FaissSearchResponse{}, fmt.Errorf("faiss: error searching index: %w", err)
	}
	return FaissSearchResponse{
		Distances: distances,
		Labels:    labels,
	}, nil
}

func (r *FaissClient) exists(indexName string) bool {
	path := filepath.Join(r.dir, indexName)
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (r *FaissClient) write(index *faiss.IndexImpl, indexName string) error {
	indexPath := filepath.Join(r.dir, indexName+".index")
	err := faiss.WriteIndex(index, indexPath)
	if err != nil {
		return fmt.Errorf("faiss: error saving index: %w", err)
	}
	return nil
}

func (r *FaissClient) load(indexName string) (*faiss.IndexImpl, error) {
	exists := r.exists(indexName)
	if !exists {
		return nil, fmt.Errorf("faiss: index %s does not exist", indexName)
	}

	indexPath := filepath.Join(r.dir, indexName+".index")

	// Load the index from disk
	index, err := faiss.ReadIndex(indexPath, 0)
	if err != nil {
		return nil, fmt.Errorf("faiss: failed to load index from %s: %w", indexPath, err)
	}

	return index, nil
}
