package api

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/santacodes/SecureEx/server/api/stats"
	"github.com/santacodes/SecureEx/server/database"
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
	isCached, isSafe := database.AlreadyCached(domain)
	if isCached {
		log.Println(domain + " is " + strconv.FormatBool(isSafe))
		// return isSafe
		return
	}

	// check for ssl verification
	// check for cloudflare

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
	var flag int
	var domaindata JSONdata
	check := json.Valid(jsondata)
	if check {
		log.Println("Checked and Valid Data!")
		json.Unmarshal(jsondata, &domaindata)
		log.Println("parsed to JSON")
		log.Println("Domain age for", domain, "is", domaindata.DomainAge, "days")

		//Check if the domain is using cloudflare or google dns
		if strings.Contains(string(domaindata.Status), "cloudflare") || strings.Contains(string(domaindata.Status), "google") {
			flag = 0
		} else {
			flag = 1
		}
		if checkSSLCertificate(domain) {
			database.Append(domain, true)
		} else {
			stats.Calc(1, domaindata.DomainAge, flag) //have to edit the 3rd parameter
		}
	} else {
		log.Println("Invalid Data")
	}

}

// check for SSL certificate
func checkSSLCertificate(domain string) bool {
	log.Println("Checking for SSL")
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		log.Println("Server doesn't support SSL certificate err: " + err.Error())
		return false
	} else {
		log.Println("Host has SSL Certificate")
		err = conn.VerifyHostname(domain)
		if err != nil {
			log.Println("Hostname doesn't match with certificate: " + err.Error())
			return false
		} else {
			log.Println("Hosts name matches with SSL ")
			return true
		}
	}
}
