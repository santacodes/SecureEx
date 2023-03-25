package stats

type weight struct {
	Domain float64
	Status float64
	Age    float64
}

func calc() float64 {
	var node = new(weight)
	node.Domain = 0.5
	node.Status = 0.3
	node.Age = 0.6

	return 1 //The amount of safety

}
