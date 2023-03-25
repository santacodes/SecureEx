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

func calc(domain string, age int) float64 {
	var node = new(weight)
	node.Domain = 0.5
	node.Status = 0.3
	node.Age = 0.6

	//Get all the values in boolean and multiply them with the weights and add the results to get the predicted value

	if strings.Contains(domain, "https") {
		fmt.Println("Safe")
	} else {
		fmt.Println("unsafe")
	}

	if age > 1000 {
		fmt.Println("safe")
	} else {
		fmt.Println("Unsafe")
	}

	return 1 //The amount of safety

}
