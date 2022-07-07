package main

const (
	dogFoodPerDay = float64(10) / (5 * 30)
	catFoodPerDay = float64(7) / (1 * 30)
	cowFoodPerDay = float64(25) / (1 * 30)
)

type dog struct {
	weight float64
	
}

func (d dog) getName() string {
	return "dog"
}
func (d dog) getWeight() float64 {
	return d.weight
}
func (d dog) getFoodPerDay() float64 {
	return d.weight*dogFoodPerDay
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
func (c cat) getFoodPerDay() float64 {
	return c.weight*catFoodPerDay
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
func (c cow) getFoodPerDay() float64 {
	return c.weight*cowFoodPerDay
}


