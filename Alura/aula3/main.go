package main

import (
	"aula3/roots"
	"net/http"
)

func main() {
	roots.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
