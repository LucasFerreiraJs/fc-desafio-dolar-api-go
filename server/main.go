package main

import (
	"cotacaoModulo/db"
	"cotacaoModulo/request"
	"encoding/json"
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

	w.Header().Set("Content-Type", "application/json")

	resBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao formatar response %v\n", err)
	}

	w.Write([]byte(resBytes))
	print("res")
	print("\n")
	fmt.Println(res)

	print("\n")
	print("Registrando...")
	print("\n")

	db.RegisterDolarDB(ctx, res.Bid)
}
