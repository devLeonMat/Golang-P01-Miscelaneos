package main

import "fmt"

type human interface {
	breathe()
	think()
	eat()
	sex() string
}
type animal interface {
	breathe()
	eat()
	isCarnivore() bool
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
}
type woman struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
}

func (h *man) breathe()    { h.breathing = true }
func (h *man) eat()        { h.eating = true }
func (h *man) think()      { h.thinking = true }
func (h *man) sex() string { return "man" }

func (w *woman) breathe()    { w.breathing = true }
func (w *woman) eat()        { w.eating = true }
func (w *woman) think()      { w.thinking = true }
func (w *woman) sex() string { return "woman" }

func HumanBreath(hu human) {
	hu.breathe()
	fmt.Printf("I'm a %s, and I'm already breathing \n", hu.sex())
}

func main() {
	Pedro := new(man)
	HumanBreath(Pedro)

	Maria := new(woman)
	HumanBreath(Maria)

}
