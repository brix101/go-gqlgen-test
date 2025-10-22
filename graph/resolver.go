package graph

import "github.com/brix101/gqlgen-todos/todo"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoStore *todo.MemoryStore
}
