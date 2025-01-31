package game

import (
	"fmt"
	"testing"
)

type IAnimal interface {
	Speak()
}

type Animal struct {
	Name string
}

func (a *Animal) Init(name string) {
	a.Name = name
}

func (a *Animal) Do(animal IAnimal) {
	animal.Speak()
}

func (a *Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
	Animal // 繼承 Animal
}

func (d *Dog) Speak() {
	d.Animal.Speak()
	fmt.Println(d.Name, "says Woof!")
}

func TestGame_Time(t *testing.T) {
	animal := &Dog{}
	animal.Init("BOBO")
	animal.Speak() // Animal makes a sound
}
