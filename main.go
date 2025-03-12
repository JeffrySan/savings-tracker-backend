package main

import (
	"log"
	"time"

	helpers "github.com/JeffrySan/go-packages-helpers/packages"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	time.Sleep(4 * time.Second)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	log.Println("Calculating...")

	go CalculateValue(intChan)
	log.Println("Continuing...")
	num := <-intChan
	log.Println(num)
	log.Println("Resulting...")
}
