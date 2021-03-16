package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
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


func main() {
	baseUrl := "https://torgi.gov.ru/lotSearch1.html"
	searchType := "?bidKindId=8"

	userAgent, err := GetRandom(UserAgents)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userAgent)

	//collyParse(url, userAgent)



	_, err = httpParse(baseUrl, searchType, userAgent)
	if err != nil {
		fmt.Println(err)
	}

	return
}
// id11_hf_0=&common%3AbidOrganizationName=&common%3AbidNumber=&common%3ApropertyTypes%3AmultiSelectText=&common%3AbidFormId=&common%3Acountry=185&common%3AlocationKladrCommon=&common%3AkladrIdStr=&search_panel%3AbuttonsPanel%3Asearch=1
// id11_hf_0:
// common:bidOrganizationName:
// common:bidNumber:
// common:propertyTypes:multiSelectText:
// common:bidFormId:
// common:country: 185
// common:locationKladrCommon:
// common:kladrIdStr:
// search_panel:buttonsPanel:search: 1

func collyParse(url string, userAgent string) ([]byte, error) {
	collector := colly.NewCollector()
	collector.UserAgent = userAgent
	collector.Async = true

	err := collector.Visit(url)
	if err != nil {
		return nil, err
	}
	//collector.OnHTML()
	fmt.Println(collector.String())
	return nil, nil
}

func httpParse(baseUrl string, searchType string, userAgent string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl + searchType, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("user-agent", userAgent)


	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	index := bytes.Index(content, []byte("title=\"Перейти на одну страницу вперед\""))
	reg := regexp.MustCompile("href=\".*\"\\s")
	link := reg.Find(content[index:])
	if link == nil {
		return nil, fmt.Errorf("Can't find link!")
	}

	fmt.Println(string(link))
	fmt.Println(string(link[len("href=\""): len(link) - 2]))
	link = link[len("href=\""): len(link) - 2]
	req, err = http.NewRequest("GET", baseUrl + string(link), nil)
	cookies := resp.Cookies()
	fmt.Println(cookies)
	for _, cookie := range cookies {
		if cookie.Name == "JSESSIONID" {
			req.AddCookie(cookie)
			break
		}
	}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}


	return nil, nil
}
