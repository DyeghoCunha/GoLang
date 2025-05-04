package db

import "database/sql"
import _ "github.com/lib/pq"

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja  password=17122014 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
