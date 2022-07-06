package main

type dog struct {
	weight float64
	foodPerDay float64
	
}

func (d dog) getName() string {
	return "dog"
}
func (d dog) getWeight() float64 {
	return d.weight
}
func (d dog) getFoodPerDay() float64 {
	return d.foodPerDay
}

type cat struct {
	weight float64
	foodPerDay float64
}

func (c cat) getName() string {
	return "cat"
}
func (c cat) getWeight() float64 {
	return c.weight
}
func (c cat) getFoodPerDay() float64 {
	return c.foodPerDay
}

type cow struct {
	weight float64
	foodPerDay float64
}

func (c cow) getName() string {
	return "cow"
}
func (c cow) getWeight() float64 {
	return c.weight
}
func (c cow) getFoodPerDay() float64 {
	return c.foodPerDay
}


