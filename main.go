package main

import (
	"alura/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8004", nil)
}
