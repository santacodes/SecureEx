package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/santacodes/SecureEx/server/api/stats"
)

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
	// check if website already cached in database

	url := ("https://api.ip2whois.com/v2?key=96C50BC55507EAD854520B88AA6C55F8&domain=" + domain)
	log.Println("Your Domain is", domain)

	res, err := http.Get(url)
	if err != nil {
		log.Println("Error making http request to ip2whois api", err)
	}
	log.Println("Received response from ip2whois")
	if res.StatusCode != 200 {
		log.Println(domain, "does not exist")
		// return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	log.Println(string(body))

	jsondata := []byte(body)

	var domaindata JSONdata
	check := json.Valid(jsondata)
	if check {
		log.Println("Checked and Valid Data!")
		json.Unmarshal(jsondata, &domaindata)
		log.Println("parsed to JSON")
		log.Println("Domain age for", domain, "is", domaindata.DomainAge, "days")
		stats.Calc(domain, domaindata.DomainAge)
	} else {
		log.Println("Invalid Data")
	}
}

// check for SSL certificate
func checkSSLCertificate(domain string) {
	fmt.Println("Checking for SSL")
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		fmt.Println("Server doesn't support SSL certificate err: " + err.Error())
	} else {
		fmt.Println("Host has SSL Certificate")
		err = conn.VerifyHostname(domain)
		if err != nil {
			fmt.Println("Hostname doesn't match with certificate: " + err.Error())
		} else {
			fmt.Println("Hosts name matches with SSL ")
		}
	}
}
