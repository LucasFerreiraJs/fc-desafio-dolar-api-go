package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cotacao struct {
	USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {
	var apiUrl = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	print(apiUrl)
	print("\n")

	req, err := http.Get(apiUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição %v\n", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta %v\n", err)
	}

	var cotacao Cotacao
	err = json.Unmarshal(res, &cotacao)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da informação %v\n", err)
	}
	fmt.Println(cotacao)

}
