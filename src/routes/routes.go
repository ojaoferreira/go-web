package routes

import (
	"github.com/ojaoferreira/goweb/src/controllers"
	"net/http"
)

func Carregar() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
}
