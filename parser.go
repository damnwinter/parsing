package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var UserAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.9 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"
}

func GetRandom(data []string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("\"data\" is empty")
	}
	rand.Seed(time.Now().Unix())
	return data[rand.Intn(len(data))], nil
}


func main() {
	_ = "https://torgi.gov.ru/lotSearch1.html?bidKindId=8"

	userAgent, err := GetRandom(UserAgents)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userAgent)

	//collyParse(url, userAgent)


	req, err := http.Get("https://torgi.gov.ru/?wicket:interface=:0:search_panel:resultTable:list:bottomToolbars:2:toolbar:span:navigator:navigation:9:pageLink::ILinkListener::")
	cont, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(cont))



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

//?wicket:interface=:2:search_panel:resultTable:list:bottomToolbars:2:toolbar:span:navigator:navigation:1:pageLink::ILinkListener::" id="id162" onclick="showBusysign();var wcall=wicketAjaxGet('?wicket:interface=:2:search_panel:resultTable:list:bottomToolbars:2:toolbar:span:navigator:navigation:1:pageLink::IBehaviorListener:0:-1',function() { }.bind(this),function() { }.bind(this), function() {return Wicket.$('id162') != null;}.bind(this));return !wcall;
// https://torgi.gov.ru/lotSearch1.html?wicket:interface=:11:search_panel:resultTable:list:bottomToolbars:2:toolbar:span:navigator:navigation:6:pageLink::IBehaviorListener:0:-1&random=0.5868151978315934
//https://torgi.gov.ru/lotSearch1.html?wicket:interface=:12:search_panel:resultTable:list:bottomToolbars:2:toolbar:span:navigator:navigation:4:pageLink::IBehaviorListener:0:-1&random=0.6980142527742021

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

func httpParse(url string, userAgent string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
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
	return content, nil
}
