package main

import "fmt"

type animals interface {
	getName() string
	getWeight() float64
}
type weightAnimal interface {
	getWeight() float64
}

func calcFoodPerDay(animalsFodPerDay []foodPerDay, name string) float64 {
	for _, v := range animalsFodPerDay {
		if v.name == name {
			return v.foodWeight / (v.normWeight * float64(v.amountDays))
		}
	}
	return 0
}

func calculationFarmFood(animals []animals, animalsFodPerDay []foodPerDay, days int) float64 {
	var foodWeight float64
	for _, v := range animals {
		foodWeightAnimal := calcFoodPerDay(animalsFodPerDay, v.getName()) * v.getWeight() * float64(days)
		fmt.Printf("Animal %s weighing %.2f kg needs %.2f kg of food fo %d days\n", v.getName(), v.getWeight(), foodWeightAnimal, days)
		foodWeight += foodWeightAnimal
	}
	return foodWeight
}

func main() {
	daysAmount := 30
	animalsFodPerDay := []foodPerDay{
		{
			name:       "dog",
			foodWeight: 10,
			normWeight: 5,
			amountDays: 30,
		},
		{
			name:       "cat",
			foodWeight: 7,
			normWeight: 1,
			amountDays: 30,
		},
		{
			name:       "cow",
			foodWeight: 25,
			normWeight: 1,
			amountDays: 30,
		},
	}

	animalsInFarm := []animals{
		dog{weight: 12},
		dog{weight: 6},
		cat{weight: 4},
		cow{weight: 450},
	}

	fmt.Println(calculationFarmFood(animalsInFarm, animalsFodPerDay, daysAmount))

}
