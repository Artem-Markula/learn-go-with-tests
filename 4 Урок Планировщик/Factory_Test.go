package main

import "testing"
import "fmt"

func TestFactoryMethod(t *testing.T) {
	fmt.Println("Test: Availability Of Guitars in Storage")

	id := []string{"Washburn", "Dean", "BcReach", "Fender", "Stratocaster", "Yamaha"}

	factory := Search()
	products := []Guitar{
		factory.CreateProduct("Washburn"),
		factory.CreateProduct("Dean"),
		factory.CreateProduct("BcReach"),
		factory.CreateProduct("Fender"),
		factory.CreateProduct("Stratocaster"),
		factory.CreateProduct("Yamaha"),
	}
	
	

	for i, product := range products {
		if action := product.Use(); action != id[i] {
			t.Errorf("Expect action to %s, but %s.\n", id[i], action)
		}
	}
}
