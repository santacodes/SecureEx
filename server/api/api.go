package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetInfo(domain string) {

	url := ("https://api.ip2whois.com/v2?key=15EDAD6CFD6CC07185515EDD2364FABC&domain=" + domain)
	fmt.Println("Your Domain is ", domain)
	fmt.Println("\n")
	fmt.Println("\n")
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
