package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
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

type LotContent struct {

}

type Lot struct {
	mainLink 	[]byte
	printLink 	[]byte
	content 	LotContent
}

type Page struct {
	url		string
	content []byte
}

func (p Page) GetLots() []Lot {
	getLotLinks(p.content, p.url)


	return nil
}


func main() {

	baseUrl := "https://torgi.gov.ru/"
	searchLink :="lotSearch1.html"
	searchType := "?bidKindId=8"

	userAgent, err := GetRandom(UserAgents)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(userAgent)


	_, err = httpParse(baseUrl, searchLink, searchType, userAgent)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func httpParse(baseUrl string, searchLink string, searchType string, userAgent string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl + searchLink + searchType, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("user-agent", userAgent)


	_ = make([][]byte, 1000)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	allLinks := make([][]byte, 1000)
	for page_num := 0;; page_num += 1 {
		currentPageLinks := getLotLinks(content, baseUrl)
		allLinks = append(allLinks, currentPageLinks...)
		fmt.Println(page_num, len(allLinks))
		//for _, link := range currentPageLinks {
		//	fmt.Println(string(link))
		//}

		index := bytes.Index(content, []byte("title=\"Перейти на одну страницу вперед\""))
		if index == -1 {
			break
		}
		reg := regexp.MustCompile("\\?wicket.*\"\\s")
		link := reg.Find(content[index:])
		if link == nil {
			return nil, fmt.Errorf("Can't find link!")
		}

		link = link[: len(link) - 2]
		fmt.Println(string(link))
		newLink := baseUrl + searchLink + string(link)
		fmt.Println(string(newLink))
		req, err = http.NewRequest("GET", newLink, nil)
		if err != nil {
			return nil, err
		}
		for _, cookie := range resp.Cookies() {
			req.AddCookie(cookie)
		}

		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
		content, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}



	return nil, nil
}


func getLotLinks(page []byte, baseUrl string) [][]byte {
	reg := regexp.MustCompile("<a href=.*title=\"Просмотр\">")
	tempLinks := reg.FindAll(page, -1)
	if tempLinks == nil {
		return nil
	}
	startPos := len("<a href=\"")
	endPos := len("\" title=\"Просмотр\">")
	for ind, tempLink := range tempLinks {
		tempLinks[ind] = append([]byte(baseUrl), tempLink[startPos : len(tempLink) - endPos]...)
	}
	return tempLinks
}
