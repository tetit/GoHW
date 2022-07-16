package main

const (
	dogFoodPerDay = float64(10) / (5 * 30)
	catFoodPerDay = float64(7) / (1 * 30)
	cowFoodPerDay = float64(25) / (1 * 30)
	minDodWheight = float64(2)
	minCatWheight = float64(1)
	minCowWheight = float64(350)
)

type dog struct {
	name   string
	weight float64
	edible bool
}

func (d dog) String() string {
	return d.name
}
func (d dog) getWeight() float64 {
	return d.weight
}
func (d dog) getEdible() bool {
	return d.edible
}
func (d dog) getFoodPerDay() float64 {
	return d.weight * dogFoodPerDay
}

type cat struct {
	name   string
	weight float64
	edible bool
}

func (c cat) String() string {
	return c.name
}
func (c cat) getWeight() float64 {
	return c.weight
}
func (c cat) getEdible() bool {
	return c.edible
}
func (c cat) getFoodPerDay() float64 {
	return c.weight * catFoodPerDay
}

type cow struct {
	name   string
	weight float64
	edible bool
}

func (c cow) String() string {
	return c.name
}
func (c cow) getWeight() float64 {
	return c.weight
}
func (c cow) getEdible() bool {
	return c.edible
}
func (c cow) getFoodPerDay() float64 {
	return c.weight * cowFoodPerDay
}
