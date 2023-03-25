package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	domain      Domain
	registrant  Registrant
	admin       Admin
	tech        Tech
	billing     Billing
	nameservers string
}

type Domain struct {
	domain       string `json:domain`
	domain_id    string
	status       string
	create_date  string
	update_date  string
	expire_date  string
	domain_age   int
	whois_server string
	registrar    []Registrar
}

type Registrar struct {
	iana_id int
	name    string
	url     string
}

type Registrant struct {
	name           string
	organization   string
	street_address string
	city           string
	region         string
	zip_code       string
	country        string
	phone          string
	fax            string
	email          string
}

type Admin struct {
	name           string
	organization   string
	street_address string
	city           string
	region         string
	zip_code       string
	country        string
	phone          string
	fax            string
	email          string
}
type Tech struct {
	name           string
	organization   string
	street_address string
	city           string
	region         string
	zip_code       string
	country        string
	phone          string
	fax            string
	email          string
}

type Billing struct {
	name           string
	organization   string
	street_address string
	city           string
	region         string
	zip_code       string
	country        string
	phone          string
	fax            string
	email          string
}

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
		fmt.Println("this is the domain age: " + domaindata.domain.status)
	} else {
		fmt.Println("Invalid Data")
	}
}
