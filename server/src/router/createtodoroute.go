package router

import (
	"fmt"
	"net/http"

	"github.com/drabinowitz/ChartHopInterview/server/src/businessentity"
	"github.com/drabinowitz/ChartHopInterview/server/src/service"
	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Todo businessentity.Todo `json:"todo"`
}

type CreateTodoResponse struct {
	Todo businessentity.Todo `json:"todo"`
}

type CreateTodoRouteDependencies struct {
	TodoService service.TodoService
}

func NewCreateTodoRoute(deps CreateTodoRouteDependencies) route {
	return &createTodoRoute{
		todoService: deps.TodoService,
	}
}

type createTodoRoute struct {
	todoService service.TodoService
}

func (createTodoRoute *createTodoRoute) Handle(c *gin.Context) {
	req := CreateTodoRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	todo, err := createTodoRoute.todoService.Create(req.Todo)

	if err != nil {
		fmt.Println("failed to create todo, err", err)

		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(201, CreateTodoResponse{
		Todo: *todo,
	})
}
