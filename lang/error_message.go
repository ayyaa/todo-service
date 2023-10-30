package lang

import (
	errors "github.com/ayyaa/todo-services/lib/customerrors"
)

var (
	ErrNoRowsUpdated = errors.New("No rows updated")

	ErrListNotFound    = errors.NewNotFoundError("List not found")
	ErrSubListNotFound = errors.NewNotFoundError("Sublist not found")
)
