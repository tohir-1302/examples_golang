package main

import "fmt"

type Human struct {
	Name    string
	Age     int
	Surname string
}

func NewHuman() *Human {
	return &Human{
		Name:    "Human",
		Age:     20,
		Surname: "James",
	}
}

func (h *Human) Print() {
	fmt.Printf("Name:%v\n", h.Name)
	fmt.Printf("Age:%v\n", h.Age)
	fmt.Printf("Surname:%v\n", h.Surname)
}

func (h *Human) changeAge(newAge int) {
	h.Age = newAge
	fmt.Println("Age updated successfully!!!")
}

func main() {
	newHuman := NewHuman()
	newHuman.changeAge(45)
	newHuman.Print()
}
