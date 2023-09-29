package main

import "fmt"

type Grower interface {
	growUp()
}

type person struct {
	age    int
	height int
	weight int
	male   string
	name   string
}

func NewPerson(age, height, weight int, male, name string) *person {
	return &person{
		age:    age,
		height: height,
		weight: weight,
		male:   male,
		name:   name,
	}
}

func (p *person) myNameIs() {
	fmt.Printf("My name is %s\n", p.name)
}

func (p *person) growUp() {
	p.height++
}

func (p *person) gainWeight() {
	p.weight++
}

type workman struct {
	person
	placeOfWork string
}

type dog struct {
}

func (d *dog) growUp() {
	fmt.Printf("Gav! Gav!") // just example
}

func NewWorkman(placeOfWork string, p *person) *workman {
	return &workman{
		placeOfWork: placeOfWork,
		person:      *p,
	}
}

func main() {
	p := NewPerson(30, 183, 76, "male", "Michael")
	w := NewWorkman("Factory", p)
	w.myNameIs()
	d := &dog{}

	growers := []Grower{p, d}
	for _, gr := range growers {
		gr.growUp() // polymorphism
	}
}
