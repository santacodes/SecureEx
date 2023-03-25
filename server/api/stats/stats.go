package stats

import (
	"fmt"
	"strings"
)

type weight struct {
	Domain float64
	Status float64
	Age    float64
}

type vulnerability struct {
	Domain float64
	Status float64
	Age    float64
}

func Calc(domain string, age int) float64 {
	var node = new(weight)
	node.Domain = 0.5
	node.Status = 0.1
	node.Age = 0.4

	//Get all the values in boolean and multiply them with the weights and add the results to get the predicted value

	var v1 = new(vulnerability) //init a node

	if strings.Contains(domain, "https") {
		v1.Domain = 0 * (node.Domain)
	} else {
		v1.Domain = 1 * (node.Domain)
	}
	if age > 1000 {
		v1.Age = 0 * (node.Age)
	} else {
		v1.Age = 1 * (node.Age)
	}

	var vuln float64 = v1.Domain + v1.Age + v1.Status
	fmt.Println("\n")

	fmt.Print("Vulnerability Possibilty is: ")
	fmt.Print(vuln)
	fmt.Println("\n")
	return vuln //The amount of safety

}
