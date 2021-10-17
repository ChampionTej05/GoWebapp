package main

import (
	"fmt"
	"net/http"

	"github.com/ChampionTej05/GoWebapp/controllers"
)

func main() {
	fmt.Println("Hello world")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
