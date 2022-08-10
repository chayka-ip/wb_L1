package main

import "fmt"

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры
	Human (аналог наследования).
*/

type Human struct {
	Name   string
	Age    uint8
	Salary uint64
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetAge(age uint8) {
	h.Age = age
}

func (h *Human) SayAge() {
	fmt.Printf("%s: I'm %d.\n", h.Name, h.Age)
}

func (h *Human) SayHello() {
	fmt.Printf("Hi, i'm %s.\n", h.Name)
}

func (h *Human) LiveYear() {
	h.Age++
	fmt.Printf("%s: Another year is gone... I'm %d now...\n", h.Name, h.Age)
}

type Action struct {
	// Embedding is accomplished by composition
	Human
	H      Human
	Name   string
	Age    uint64
	Result bool
}

func (a *Action) SetName(name string) {
	a.Name = name
}

func (a *Action) SayAge() {
	fmt.Printf("%s: %d years.\n", a.Name, a.Age)
}

func (a *Action) SetAge(age uint64) {
	a.Age = age
}

func main() {
	a := Action{
		Human{
			Name:   "Embedded Daniel",
			Age:    5,
			Salary: 10,
		},
		Human{
			Name:   "Jack",
			Age:    16,
			Salary: 200,
		},
		"Action",
		10000,
		true,
	}

	/*
		  Output:
		- Embedded Daniel: Another year is gone... I'm 6 now...

		Action has embeded Human type, so Human's methods and fields are available.
		LiveYear will modify data that belongs to embedded struct
		when there are now collisions or shadowing
	*/
	a.LiveYear()

	/*
		Output: 10
		Salary field belongs to Human type and does not exist on top level.
	*/
	fmt.Println(a.Salary)

	/*
		  Output:
		- Action: 10000 years.
		Top level method SayAge is shadowing embedded method SayAge from Human.
	*/
	a.SayAge()

	/*
		Changing age of Action.
		Age property belongs to Action because it exists at top level.
	*/
	a.SetAge(a.Age - 1)
	/*
		Changing age of embedded Human.
		Age property belongs to embedded Human because it is requested directly.
	*/
	a.Human.SetAge(a.Human.Age + 2)

	/*
		  Output:
		- Action: 9999 years.
		- Embedded Daniel: I'm 8.

		We can modify shadowed properties and call their methods
		by referencing them directly
	*/
	a.SayAge()
	a.Human.SayAge()

	/*
		  Output:
		- Hi, i'm Jack.
		- Hi, i'm Martin.

		Struct can be a member of other struct.
		We should reference to it directly to modify its properties and use methods.
	*/
	a.H.SayHello()
	a.H.SetName("Martin")
	a.H.SayHello()
}
