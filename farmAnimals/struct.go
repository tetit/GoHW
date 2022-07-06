package main

type dog struct {
	weight float64
}

func (d dog) getName() string {
	return "dog"
}
func (d dog) getWeight() float64 {
	return d.weight
}

type cat struct {
	weight float64
}

func (c cat) getName() string {
	return "cat"
}
func (c cat) getWeight() float64 {
	return c.weight
}

type cow struct {
	weight float64
}

func (c cow) getName() string {
	return "cow"
}
func (c cow) getWeight() float64 {
	return c.weight
}

type foodPerDay struct {
	name       string
	foodWeight float64
	normWeight float64
	amountDays int
}

func (f foodPerDay) getName() string {
	return f.name
}
