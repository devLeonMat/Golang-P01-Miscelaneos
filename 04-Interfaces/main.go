package main

import "fmt"

type livingBeing interface {
	alive() bool
}

type human interface {
	breathe()
	think()
	eat()
	sex() string
	alive() bool
}
type animal interface {
	breathe()
	eat()
	isCarnivore() bool
	alive() bool
}

type vegetable interface {
	plantClassification() string
}

//human gender
type man struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	isMan     bool
	living    bool
}
type woman struct {
	man
}

func (h *man) breathe() { h.breathing = true }
func (h *man) eat()     { h.eating = true }
func (h *man) think()   { h.thinking = true }
func (h *man) sex() string {
	if h.isMan == true {
		return "man"
	} else {
		return "woman"
	}
}
func (h *man) alive() bool { return h.living }

func HumanBreath(hu human) {
	hu.breathe()
	fmt.Printf("I'm a %s, and I'm already breathing \n", hu.sex())
}

/*************** ANIMAL ***/
type dog struct {
	breathing bool
	eating    bool
	carnivore bool
	living    bool
}

func (d *dog) breathe()          { d.breathing = true }
func (d *dog) eat()              { d.eating = true }
func (d *dog) isCarnivore() bool { return d.carnivore }
func (d *dog) alive() bool       { return d.living }

func AnimalBreath(an animal) {
	an.breathe()
	fmt.Println("I'm a animal and I'm already breathing")
}

func AnimalCarnivore(an animal) int {
	if an.isCarnivore() == true {
		return 1
	}
	return 0
}

func IAmAlive(s livingBeing) bool {
	return s.alive()
}

func main() {
	totalCarnivores := 0
	Dogo := new(dog)
	Dogo.carnivore = true
	Dogo.living = true
	AnimalBreath(Dogo)
	totalCarnivores = +AnimalCarnivore(Dogo)

	fmt.Printf("Total carnivores %d \n", totalCarnivores)
	fmt.Printf("I am alive %t", IAmAlive(Dogo))
	//Pedro := new(man)
	//Pedro.isMan = true
	//HumanBreath(Pedro)
	//
	//Maria := new(woman)
	//HumanBreath(Maria)

}
