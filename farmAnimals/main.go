package main

import "fmt"

type animals interface {
	getName() string
	getWeight() float64
	getFoodPerDay() float64
}

func calculationFarmFood(animals []animals, days int) float64 {
	var foodWeight float64
	for _, v := range animals {
		foodWeightAnimal := v.getFoodPerDay() * v.getWeight() * float64(days)
		fmt.Printf("Animal %s weighing %.2f kg needs %.2f kg of food fo %d days\n", v.getName(), v.getWeight(), foodWeightAnimal, days)
		foodWeight += foodWeightAnimal
	}
	return foodWeight
}

func main() {
	daysAmount := 30
	
	animalsInFarm := []animals{
		dog{weight: 12, foodPerDay: float64(10)/(5*30)},
		dog{weight: 6, foodPerDay: float64(10)/(5*30)},
		cat{weight: 4, foodPerDay: float64(7)/30},
		cow{weight: 450, foodPerDay: float64(25)/30},
	}

	fmt.Println(calculationFarmFood(animalsInFarm, daysAmount))

}
