//go:generate go run github.com/99designs/gqlgen@v0.14.0 generate
package graph

import "github.com/kashee337/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}
