package main

import (
	"github.com/drabinowitz/ChartHopInterview/server/src/router"
)

func main() {
	r := router.New()

	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
