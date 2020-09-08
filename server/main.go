package main

import (
	"github.com/drabinowitz/ChartHopInterview/server/src/router"
	"github.com/drabinowitz/ChartHopInterview/server/src/service"
)

func main() {
	todoService, err := service.NewTodoService()

	if err != nil {
		panic(err.Error())
	}

	r := router.New(router.Dependencies{
		TodoService: todoService,
	})

	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
