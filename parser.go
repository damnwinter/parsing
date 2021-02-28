package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var UserAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.9 Safari/537.36",
}

func GetRandom(data []string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("\"data\" is empty")
	}
	rand.Seed(time.Now().Unix())
	return data[rand.Intn(len(data))], nil
}


func main() {
	url := "https://torgi.gov.ru/lotSearch1.html?bidKindId=8"

	userAgent, err := GetRandom(UserAgents)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userAgent)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("user-agent", userAgent)


	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))

	return
}
