package main

import (
	"fmt"
	"github.com/ojaoferreira/goweb/src/routes"
	"net/http"
)

func main() {
	fmt.Println("Starting server on 8000")

	routes.Carregar()

	panic(http.ListenAndServe(":8000", nil))
}
