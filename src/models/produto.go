package models

import (
	"github.com/ojaoferreira/goweb/src/db"
	"log"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.Conectar()
	defer db.Close()

	p := Produto{}
	produtos := []Produto{}

	rows, err := db.Query("select * from produtos order by nome")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade); err != nil {
			panic(err)
		}
		produtos = append(produtos, p)
	}

	return produtos
}

func CriaNovoProduto(p Produto) {
	db := db.Conectar()
	defer db.Close()

	stmt, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		log.Println("Error ao inserir um produto no banco:", err)
	}

	stmt.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade)
	defer stmt.Close()
}

func DeletaProduto(id int) {
	db := db.Conectar()
	defer db.Close()

	stmt, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		log.Println("Erro ao buscar um produto:", err)
	}
	defer stmt.Close()

	stmt.Exec(id)
}

func BuscarProdutoPorId(id int) Produto {
	var p Produto

	db := db.Conectar()
	defer db.Close()

	row := db.QueryRow("select * from produtos where id=$1", id)
	err := row.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
	if err != nil {
		log.Println(err)
	}

	return p
}

func AtualizaProduto(p Produto) {
	db := db.Conectar()
	defer db.Close()

	log.Println(p)

	stmt, _ := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	res, _ := stmt.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade, p.Id)

	rowsAffected, _ := res.RowsAffected()
	log.Println("RowsAffected:", rowsAffected)
}
