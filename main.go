package main

import (
	"github.com/hakimehmordadi/GOlang-Restful/api/view"
)

func main() {

	// Set up mode to release
	// gin.SetMode(gin.ReleaseMode)

	// Where everything is going to start
	view.StartServer()
}

