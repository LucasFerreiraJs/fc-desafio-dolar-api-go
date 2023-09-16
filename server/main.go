package main

import (
	"cotacaoModulo/db"
	"cotacaoModulo/register"
	"cotacaoModulo/request"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	print("on \n")
	db.Connection()

	http.HandleFunc("/cotacao", handleCotacao)
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func handleCotacao(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	defer log.Println("fim request")
	var apiUrl = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	res, err := request.RequestDollarValue(ctx, apiUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição %v\n", err)
	}
	print("\n")
	print("Registrando...")
	print("\n")
	fmt.Printf(res.Name)
	print("\n")

	db.RegisterDolarDB(ctx, res.Bid)
	register.SaveToFile(res.Bid)
}
