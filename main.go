package main

import "fmt"

func main() {
	fmt.Println("something")

	var madyar Person

	madyar.printName()

}

type Person struct {
	Name   string
	Age    int
	Height int
	Weight int
}

func (p *Person) printName() {
	fmt.Println(p.Name)
}

// обьявление мапы
// Посчитать кол-во элементов и положить в мапе
// пройтись по мапе и в случае если есть кол-во элементов больше 1 то return false