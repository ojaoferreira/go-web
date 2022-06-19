package controllers

import (
	"github.com/ojaoferreira/goweb/src/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	temp = template.Must(template.ParseGlob("src/templates/*.html"))
)

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	var p models.Produto
	var err error

	if r.Method == "POST" {
		p.Id, _ = strconv.Atoi(r.FormValue("id"))
		p.Nome = r.FormValue("nome")
		p.Descricao = r.FormValue("descricao")
		if p.Preco, err = strconv.ParseFloat(r.FormValue("preco"), 64); err != nil {
			log.Println("Erro ao converter preco:", err)
		}
		if p.Quantidade, err = strconv.Atoi(r.FormValue("quantidade")); err != nil {
			log.Println("Erro ao converter quantidade:", err)
		}

		if p.Id == 0 {
			models.CriaNovoProduto(p)
		} else {
			models.AtualizaProduto(p)
		}
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Erro ao converter id:", err)
	}
	models.DeletaProduto(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	produto := models.BuscarProdutoPorId(id)

	temp.ExecuteTemplate(w, "Edit", produto)
}
