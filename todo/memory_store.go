package todo

import (
	"sync"

	"github.com/brix101/gqlgen-todos/graph/model"
)

type MemoryStore struct {
	mu        sync.Mutex
	todoItems []*model.Todo
	subs      map[string]chan *model.Todo
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		todoItems: []*model.Todo{},
		subs:      make(map[string]chan *model.Todo),
	}
}
