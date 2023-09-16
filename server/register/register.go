package register

import (
	"fmt"
	"os"
)

func SaveToFile(value string) {

	file, err := os.Create("arquivo.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %v\n", err)
	}

	var valorAtual = fmt.Sprintf("Dolar : %v", value)
	print(valorAtual)
	_, err = file.WriteString(valorAtual)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao registrar informação no arquivo %v\n", err)
	}

	defer file.Close()

}
