package router

import (
	"time"

	"github.com/drabinowitz/ChartHopInterview/server/src/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	TodoService service.TodoService
}

// New generates a gin adapter for handling our routes
func New(deps Dependencies) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
		MaxAge:        12 * time.Hour,
	}))

	createTodoRoute := NewCreateTodoRoute(CreateTodoRouteDependencies{
		TodoService: deps.TodoService,
	})

	listTodoRoute := NewListTodoRoute(ListTodoRouteDependencies{
		TodoService: deps.TodoService,
	})

	grouping := r.Group("todos")

	grouping.POST("/", func(c *gin.Context) {
		createTodoRoute.Handle(c)
	})

	grouping.GET("/", func(c *gin.Context) {
		listTodoRoute.Handle(c)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
