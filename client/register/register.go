package register

import (
	"fmt"
	"os"
)

func SaveToFile(value string) {

	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %v\n", err)
	}

	var curentValue = fmt.Sprintf("Dolar : %v", value)
	print(curentValue)
	_, err = file.WriteString(curentValue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao registrar informação no arquivo %v\n", err)
	}

	defer file.Close()

}
