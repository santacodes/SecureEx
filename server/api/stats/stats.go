package stats

import (
	"fmt"
	"math"
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

// Sigmoid returns the input values in the range of -1 to 1

func Sigmoid(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return []float64{}, nil
	}
	s := make([]float64, len(input))
	for i, v := range input {
		s[i] = 1 / (1 + math.Exp(-v))
	}
	return s, nil
}
func Calc(domain string, age int, status string) float64 {
	//IF THIS FUNCTION IS CALLED THEN SSL IS NOT VERIFIED

	var node = new(weight)
	node.Domain = 0.6
	node.Status = 0.1
	node.Age = 0.4

	//Get all the values in boolean and multiply them with the weights and add the results to get the predicted value
	if strings.Contains("status", "cloudflare") {
		node.Status = 0
	} else {
		node.Status = 0.1
	}

	var v1 = new(vulnerability) //init a node

	if v1.Age < 200 {
		node.Domain = 0.6
		node.Age = 0.9
	} else if (v1.Age < 500) && (v1.Age >= 200) {
		node.Domain = 0.6
		node.Age = 0.7
	} else if (v1.Age >= 500) && (v1.Age < 800) {
		node.Domain = 0.4
		node.Age = 0.2
	} else {
		node.Domain = 0.4
		node.Age = 0.1
	}

	var vuln float64 = v1.Domain*node.Domain + v1.Age*node.Age + v1.Status*node.Status
	//sigout, err := Sigmoid(vuln)
	fmt.Print("Vulnerability Possibilty is: ")
	fmt.Print(vuln)
	//fmt.Println(sigout)
	//fmt.Println(err)
	return vuln //The amount of safety

}
