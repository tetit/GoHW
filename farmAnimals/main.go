package main

import (
	"errors"
	"fmt"
	"strings"
)

type animals interface {
	fmt.Stringer
	getWeight() float64
	getFoodPerDay() float64
	getEdible() bool
}

var errTypeAnimal error = errors.New("wrong name")
var errWeightAnimal error = errors.New("weight problem")
var errEdibleAnimal error = errors.New("wrong concept edible")

func validateTypeAnimal(animal animals) error {
	var animalType string
	switch animal.(type) {
	case dog:
		animalType = "dog"
	case cat:
		animalType = "cat"
	case cow:
		animalType = "cow"
	}
	if !strings.Contains(animal.String(), animalType) {
		return errTypeAnimal
	}
	return nil
}

func validateWeightAnimal(animal animals) error {
	switch animal.(type) {
	case dog:
		if animal.getWeight() < minDodWheight {
			return errWeightAnimal
		}
	case cat:
		if animal.getWeight() < minCatWheight {
			return errWeightAnimal
		}
	case cow:
		if animal.getWeight() < minCowWheight {
			return errWeightAnimal
		}
	}
	return nil
}

func validateEdibleAnimal(animal animals) error {
	switch animal.(type) {
	case dog:
		if animal.getEdible() {
			return fmt.Errorf("this animal is not edible: %w", errEdibleAnimal)
		}
	case cat:
		if animal.getEdible() {
			return fmt.Errorf("this animal is not edible: %w", errEdibleAnimal)
		}
	case cow:
		if !animal.getEdible() {
			return fmt.Errorf("this animal is edible: %w", errEdibleAnimal)
		}
	}
	return nil
}

func calculationFarmFood(animals []animals, days int) float64 {
	var foodWeight float64
	for _, v := range animals {

		if err := validateTypeAnimal(v); err != nil {
			fmt.Printf("\nanimal %s name does not match its type %T: %v\n\n", v, v, err)
			continue
		}

		if err := validateWeightAnimal(v); err != nil {
			fmt.Printf("\nneed to feed the animal %s, weight %.2f below normal: %v\n\n", v, v.getWeight(), err)
			continue
		}

		if err := validateEdibleAnimal(v); err != nil {
			fmt.Printf("\nterrible farm %s: %v\n\n", v, err)
			return 0
		}

		foodWeightAnimal := v.getFoodPerDay() * float64(days)
		fmt.Printf("Animal %s weighing %.2f kg needs %.2f kg of food fo %d days\n", v, v.getWeight(), foodWeightAnimal, days)
		foodWeight += foodWeightAnimal
	}
	return foodWeight
}

func main() {
	daysAmount := 30

	animalsInFarm := []animals{
		dog{name: "dog Boob", weight: 12, edible: false},
		dog{name: "dog Smol", weight: 1, edible: false},
		dog{name: "cat Miki", weight: 6, edible: false},
		cat{name: "cat Nice", weight: 4, edible: false},
		cow{name: "cow Bura", weight: 450, edible: false},
	}

	fmt.Println(calculationFarmFood(animalsInFarm, daysAmount))

}
