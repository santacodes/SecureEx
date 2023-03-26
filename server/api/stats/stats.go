package stats

import (
	"fmt"
	"math"
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
func Calc(domain int, age int, status int) bool {
	//IF THIS FUNCTION IS CALLED THEN SSL IS NOT VERIFIED

	var node = new(weight)
	node.Domain = 0.6
	node.Status = 0.1
	node.Age = 0.4

	//Get all the values in boolean and multiply them with the weights and add the results to get the predicted value

	var v1 = new(vulnerability) //init a node
	v1.Age = float64(age)
	v1.Domain = float64(domain)
	v1.Status = float64(status)

	if v1.Age < 200 {
		v1.Age = 0.9
		node.Domain = 0.6
		node.Age = 0.9
	} else if (v1.Age < 500) && (v1.Age >= 200) {
		v1.Age = 0.7
		node.Domain = 0.6
		node.Age = 0.7
	} else if (v1.Age >= 500) && (v1.Age < 800) {
		v1.Age = 0.55
		node.Domain = 0.4
		node.Age = 0.2
	} else {
		v1.Age = 0.2
		node.Domain = 0.4
		node.Age = 0.1
	}

	var vuln float64 = (v1.Domain)*(node.Domain) + (v1.Age)*(node.Age) + (v1.Status)*(node.Status)
	//sigout, err := Sigmoid(vuln)
	fmt.Print("Vulnerability Possibilty is: ")
	fmt.Print(vuln)
	//return bool
	if vuln < 0.5 {
		return false
	} else {
		return true
	}
}
