package request

import (
	"context"
	"cotacaoModulo/dto"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func RequestDollarValue(ctx context.Context, url string) (*dto.CotacaoDolar, error) {

	select {
	case <-time.After(200 * time.Millisecond):

		// req, err := http.NewRequestWithContext(ctx, "GET", apiUrl, nil)
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição %v\n", err)
			return nil, err
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta %v\n", err)
			return nil, err
		}

		var resultCotacao dto.CotacaoDolar
		err = json.Unmarshal(res, &resultCotacao)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da informação %v\n", err)
			return nil, err
		}
		// fmt.Println(resultCotacao)

		return &resultCotacao, nil

	case <-ctx.Done():
		log.Println("Request cancelada")
		return nil, errors.New("Request cancelada")
	}
}
