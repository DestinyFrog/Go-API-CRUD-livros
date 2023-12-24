package models

import (
	"fmt"

	"bipbop/db"
)

type Livro struct {
	Id int64
	Titulo string
	Autor string
	Capa string
	AnoDePublicacao int64
}

type LivroRequest struct {
	Titulo				string		`json:"titulo"`
	Autor				string		`json:"autor"`
	Capa				string		`json:"capa"`
	AnoDePublicacao		int64		`json:"ano_de_publicacao"`
}

func (livro Livro) ToString() string {
	return fmt.Sprintf("%d. %s (%d) - %s",
		livro.Id, livro.Titulo, livro.AnoDePublicacao ,livro.Autor)
}

func GetAll() (data []Livro, err error) {
	conn := db.OpenConnection()
	defer conn.Close()

	sql := "SELECT * FROM livros ORDER BY id"
	rows, err := conn.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var l Livro
		rows.Scan(&l.Id, &l.Titulo, &l.Autor, &l.AnoDePublicacao, &l.Capa)
		data = append(data, l)
	}

	return
}

func Insert(livro LivroRequest) (err error) {
	conn := db.OpenConnection()
	defer conn.Close()

	sql := "INSERT INTO livros ( titulo, autor, ano_de_publicacao, capa ) VALUES ($1,$2,$3,$4) RETURNING id"
	_,err = conn.Query(sql, livro.Titulo, livro.Autor, livro.AnoDePublicacao, livro.Capa)
	return
}

func Update(id int64, livro LivroRequest) (err error) {
	conn := db.OpenConnection()
	defer conn.Close()

	sql := "UPDATE livros SET titulo = $2, autor = $3, ano_de_publicacao = $4, capa = $5 WHERE id = $1"
	_, err = conn.Query(sql, id, livro.Titulo, livro.Autor, livro.AnoDePublicacao, livro.Capa)
	return
}

func Delete(id int64) (err error) {
	conn := db.OpenConnection()
	defer conn.Close()

	sql := "DELETE FROM livros WHERE id = $1"
	_, err = conn.Query(sql, id)
	return
}