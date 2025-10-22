package todo

import (
	"github.com/brix101/gqlgen-todos/graph/model"
	"github.com/google/uuid"
)

func (s *MemoryStore) Create(input model.NewTodo) *model.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := &model.Todo{
		ID:   uuid.New().String(),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "Placeholder User",
		},
	}

	s.todoItems = append(s.todoItems, todo)

	// Notify subscribers
	for _, ch := range s.subs {
		select {
		case ch <- todo:
		default:
		}
	}

	return todo
}

func (s *MemoryStore) All() []*model.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]*model.Todo{}, s.todoItems...) // copy to avoid race
}

func (s *MemoryStore) Subscribe() (<-chan *model.Todo, func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	ch := make(chan *model.Todo, 1)
	s.subs[id] = ch

	cleanup := func() {
		s.mu.Lock()
		delete(s.subs, id)
		close(ch)
		s.mu.Unlock()
	}

	return ch, cleanup
}
