package request

import (
	"context"
	"cotacaoClient/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func RequestDolarValue(ctx context.Context) (*dto.CotacaoDolar, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao configurar request %v\n", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao realizar requisição %v\n", err)
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler informações da requisição %v\n", err)
		return nil, err
	}

	var dolarValue dto.CotacaoDolar
	err = json.Unmarshal(data, &dolarValue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro json.Unmarshal  %v\n", err)
		return nil, err
	}

	return &dolarValue, nil
}
