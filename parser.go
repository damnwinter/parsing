package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://torgi.gov.ru/lotSearch1.html?bidKindId=8"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
