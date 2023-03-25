package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*

DOMAIN RETURN TYPE TEMPLATE

{
    "domain": "locaproxy.com",
    "domain_id": "1710914405_DOMAIN_COM-VRSN",
    "status": "clientTransferProhibited https://icann.org/epp#clientTransferProhibited",
    "create_date": "2012-04-03T02:34:32Z",
    "update_date": "2021-12-03T02:54:57Z",
    "expire_date": "2024-04-03T02:34:32Z",
    "domain_age": 3863,
    "whois_server": "whois.godaddy.com",
    "registrar": {
        "iana_id": "146",
        "name": "GoDaddy.com, LLC",
        "url": "https://www.godaddy.com"
    },
    "registrant": {
        "name": "Registration Private",
        "organization": "Domains By Proxy, LLC",
        "street_address": "DomainsByProxy.com",
        "city": "Tempe",
        "region": "Arizona",
        "zip_code": "85284",
        "country": "US",
        "phone": "+1.4806242599",
        "fax": "+1.4806242598",
        "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
    },
    "admin": {
        "name": "Registration Private",
        "organization": "Domains By Proxy, LLC",
        "street_address": "DomainsByProxy.com",
        "city": "Tempe",
        "region": "Arizona",
        "zip_code": "85284",
        "country": "US",
        "phone": "+1.4806242599",
        "fax": "+1.4806242598",
        "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
    },
    "tech": {
        "name": "Registration Private",
        "organization": "Domains By Proxy, LLC",
        "street_address": "DomainsByProxy.com",
        "city": "Tempe",
        "region": "Arizona",
        "zip_code": "85284",
        "country": "US",
        "phone": "+1.4806242599",
        "fax": "+1.4806242598",
        "email": "Select Contact Domain Holder link at https://www.godaddy.com/whois/results.aspx?domain=LOCAPROXY.COM"
    },
    "billing": {
        "name": "",
        "organization": "",
        "street_address": "",
        "city": "",
        "region": "",
        "zip_code": "",
        "country": "",
        "phone": "",
        "fax": "",
        "email": ""
    },
    "nameservers": "vera.ns.cloudflare.com, walt.ns.cloudflare.com"
}
*/

type JSONdata struct {
	Domain      string `json:"domain"`
	DomainId    string `json:"domain_id"`
	Status      string `json:"status"`
	CreateDate  string `json:"create_date"`
	UpdateDate  string `json:"update_date"`
	ExpireDate  string `json:"expire_date"`
	DomainAge   int    `json:"domain_age"`
	WhoIsServer string `json:"whois_server"`
	Registrar   Registrar
	Registrant  Registrant
	Admin       Admin
	Tech        Tech
	Billing     Billing
	NameServers string `json:"nameservers"`
}

type Registrar struct {
	IanaId int    `json:"iana_id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

type Registrant struct {
	Name          string `json:"name"`
	Organization  string `json:"organization"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Region        string `json:"region"`
	ZipCode       string `json:"zip_code"`
	Country       string `json:"country"`
	Phone         string `json:"phone"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
}

type Admin struct {
	Name          string `json:"name"`
	Organization  string `json:"organization"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Region        string `json:"region"`
	ZipCode       string `json:"zip_code"`
	Country       string `json:"country"`
	Phone         string `json:"phone"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
}
type Tech struct {
	Name          string `json:"name"`
	Organization  string `json:"organization"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Region        string `json:"region"`
	ZipCode       string `json:"zip_code"`
	Country       string `json:"country"`
	Phone         string `json:"phone"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
}

type Billing struct {
	Name          string `json:"name"`
	Organization  string `json:"organization"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Region        string `json:"region"`
	ZipCode       string `json:"zip_code"`
	Country       string `json:"country"`
	Phone         string `json:"phone"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
}

func GetInfo(domain string) {

	url := ("https://api.ip2whois.com/v2?key=15EDAD6CFD6CC07185515EDD2364FABC&domain=" + domain)
	fmt.Println("Your Domain is ", domain)
	fmt.Println("\n")
	fmt.Println("\n")
	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error making http request to ip2whois api", err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println("\n")
	fmt.Println(string(body))

	jsondata := []byte(body)

	var domaindata JSONdata
	check := json.Valid(jsondata)
	if check {
		fmt.Println("Checked and Valid Data! ")
		json.Unmarshal(jsondata, &domaindata)
		fmt.Println("\n")
		fmt.Printf("%#v\n", domaindata)
		fmt.Println("this is the domain age: ", domaindata.DomainAge)
	} else {
		fmt.Println("Invalid Data")
	}
}
