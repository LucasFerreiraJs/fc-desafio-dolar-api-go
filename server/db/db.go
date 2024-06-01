package db

import (
	"context"
	"cotacaoModulo/dto"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connection() {
	print("connection \n")
	var err error
	db, err = gorm.Open(sqlite.Open("dolar.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao iniciar conexão com o banco de dados: %v", err)
	}

	db.AutoMigrate(&dto.Dolar{})

}

func RegisterDolarDB(ctx context.Context, value string) {

	select {
	case <-time.After(10 * time.Millisecond):
		db.Create(&dto.Dolar{
			Valor: value,
		})

		log.Println("Valor registrado com sucesso")
	case <-ctx.Done():
		log.Println("Erro ao registrar informação no banco de dados")
	}
}
