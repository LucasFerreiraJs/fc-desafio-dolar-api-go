package main

import (
	"context"
	"cotacaoClient/register"
	"cotacaoClient/request"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	dolarValue, err := request.RequestDolarValue(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao configurar request %v\n", err)
	}

	fmt.Println(dolarValue)
	register.SaveToFile(dolarValue.Bid)
	print("\n")
}
