package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/cookies"
)

func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, data io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, data)

	if erro != nil {
		return nil, erro
	}

	cookie, _ := cookies.LerCookies(r)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cookie["token"]))

	client := &http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
