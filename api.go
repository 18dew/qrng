package qrng

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apibase = "https://qrng.anu.edu.au/API/jsonI.php"

type ApiResponse struct {
	Length int    `json:"length"`
	Data   []uint `json:"data"`
}

// queryDefaultSize will request ten numbers from the API.
func queryDefaultSize(dt datatype) ([]uint, error) {
	return queryApi(dt, 16)
}

// queryApi performs the actual request against the anu.edu servers.
func queryApi(dt datatype, size int) ([]uint, error) {
	query := fmt.Sprintf("%v?length=%d&type=%v", apibase, size, dt)
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	resp := new(ApiResponse)
	json.NewDecoder(response.Body).Decode(resp)
	fmt.Println(resp.Data, query)
	return resp.Data, nil
}
