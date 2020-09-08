package service

import (
	"fmt"
	"strconv"

	"github.com/drabinowitz/ChartHopInterview/server/src/businessentity"
)

type ListRequest struct{}

type TodoService interface {
	Create(todo businessentity.Todo) (*businessentity.Todo, error)
	Update(todo businessentity.Todo) error
	Read(userID, todoID string) (*businessentity.Todo, error)
	List(userID string, listRequest *ListRequest) ([]*businessentity.Todo, error)
}

func NewTodoService() (TodoService, error) {
	return &todoService{
		data: map[string][]*businessentity.Todo{},
	}, nil
}

type todoService struct {
	// string here is the user id
	data map[string][]*businessentity.Todo
}

func (todoService *todoService) Create(todo businessentity.Todo) (*businessentity.Todo, error) {
	existingTodos := todoService.data[todo.UserID]

	if existingTodos == nil {
		existingTodos = make([]*businessentity.Todo, 0)
	}

	todo.ID = strconv.Itoa(len(existingTodos))

	existingTodos = append(existingTodos, &todo)

	todoService.data[todo.UserID] = existingTodos

	return &todo, nil
}

func (todoService *todoService) Update(todo businessentity.Todo) error {
	existingTodos := todoService.data[todo.UserID]

	for _, existingTodo := range existingTodos {
		if existingTodo.ID == todo.ID {
			*existingTodo = todo
			return nil
		}
	}

	return fmt.Errorf("failed to fiund todo matching=[%v]", todo)
}

func (todoService *todoService) Read(userID, todoID string) (*businessentity.Todo, error) {
	existingTodos := todoService.data[userID]

	for _, existingTodo := range existingTodos {
		if existingTodo.ID == todoID {
			return existingTodo, nil
		}
	}

	return nil, fmt.Errorf("failed to fiund todo matching=[%v, todoID]", userID, todoID)
}

func (todoService *todoService) List(userID string, listRequest *ListRequest) ([]*businessentity.Todo, error) {
	existingTodos := todoService.data[userID]
	return existingTodos, nil
}
