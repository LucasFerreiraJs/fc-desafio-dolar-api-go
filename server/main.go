package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dolar struct {
	ID    int `gorm:"primaryKey"`
	Valor string
	gorm.Model
}

type CotacaoDolar struct {
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

var db *gorm.DB

func main() {
	strConn := "root:root@tcp(localhost:3306)/cotacao?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(strConn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao iniciar conexão com banco de dados %v\n", err)
	}
	db.AutoMigrate(&Dolar{})

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		defer log.Println("fim request")
		var apiUrl = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

		select {
		case <-time.After(200 * time.Millisecond):

			// req, err := http.NewRequestWithContext(ctx, "GET", apiUrl, nil)
			req, err := http.Get(apiUrl)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao fazer requisição %v\n", err)
			}
			defer req.Body.Close()

			res, err := io.ReadAll(req.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao ler resposta %v\n", err)
			}

			var cotacao CotacaoDolar
			err = json.Unmarshal(res, &cotacao)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao fazer parse da informação %v\n", err)
			}
			fmt.Println(cotacao)

			log.Println("Request processada com sucesso")
			registrarCotacao(ctx, cotacao.Bid)
			json.NewEncoder(w).Encode(cotacao)

		case <-ctx.Done():
			log.Println("Request cancelada")
		}

	})
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func registrarCotacao(ctx context.Context, valor string) {

	select {
	case <-time.After(10 * time.Millisecond):
		db.Create(&Dolar{
			Valor: valor,
		})
		fmt.Println(valor)
		log.Println("Valor registrado com sucesso")
	case <-ctx.Done():
		log.Println("Erro ao registrar informação no banco de dados")
	}

}
