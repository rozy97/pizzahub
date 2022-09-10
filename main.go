package main

import (
	"log"
	"sync"
	"time"
)

type Chef struct {
	ID   int
	Name string
}

type Product struct {
	ID   int
	Name string
}

func AddChefToKithen(kitchen chan *Chef, chef *Chef, msg string) {
	kitchen <- chef
	log.Printf("chef %s %s to kitchen", chef.Name, msg)

}

func GetChefFromKitchen(kitchen chan *Chef) *Chef {
	chef := <-kitchen
	return chef
}

func CookPizza(kitchen chan *Chef, orderID int) {
	chef := GetChefFromKitchen(kitchen)
	log.Printf("start cookin pizza from order %v", orderID)
	time.Sleep(3 * time.Second)
	AddChefToKithen(kitchen, chef, "back from cooking pizza")
	result := &Product{
		ID:   1,
		Name: "pizza",
	}
	log.Printf("%s cooked by %s!\n", result.Name, chef.Name)
}

func main() {
	var wg sync.WaitGroup
	kitchen := make(chan *Chef, 1000)
	chef1 := &Chef{
		ID:   1,
		Name: "arnold",
	}
	chef2 := &Chef{
		ID:   2,
		Name: "juna",
	}

	wg.Add(2)
	go AddChefToKithen(kitchen, chef1, "added")
	go func() {
		CookPizza(kitchen, 1)
		defer wg.Done()
	}()
	go func() {
		CookPizza(kitchen, 2)
		defer wg.Done()
	}()
	time.Sleep(2 * time.Second)
	go AddChefToKithen(kitchen, chef2, "added")
	wg.Wait()

}
