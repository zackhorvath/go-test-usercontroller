package main

import (
	"net/http"

	"github.com/zackhorvath/test-webapp/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
