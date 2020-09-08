package router

import (
	"fmt"
	"net/http"

	"github.com/drabinowitz/ChartHopInterview/server/src/businessentity"
	"github.com/drabinowitz/ChartHopInterview/server/src/service"
	"github.com/gin-gonic/gin"
)

type ListTodoRequest struct {
	UserID string `query:"user"`
}

type ListTodoResponse struct {
	Todos []businessentity.Todo `json:"todos"`
}

type ListTodoRouteDependencies struct {
	TodoService service.TodoService
}

func NewListTodoRoute(deps ListTodoRouteDependencies) route {
	return &listTodoRoute{
		todoService: deps.TodoService,
	}
}

type listTodoRoute struct {
	todoService service.TodoService
}

func (listTodoRoute *listTodoRoute) Handle(c *gin.Context) {
	req := ListTodoRequest{}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	todos, err := listTodoRoute.todoService.List(req.UserID, nil)

	if err != nil {
		fmt.Println("failed to list todo, err", err)

		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	todosToReturn := make([]businessentity.Todo, len(todos))
	for i, todo := range todos {
		todosToReturn[i] = *todo
	}

	c.JSON(201, ListTodoResponse{
		Todos: todosToReturn,
	})
}
