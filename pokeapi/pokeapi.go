package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)
func CallAPI(url string) ([]byte, error) {
	if url == "" {
		return nil, fmt.Errorf("empty url provided")
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching")
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body")
	}
	return data, nil
}