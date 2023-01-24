package internal

import "fmt"

type message string

func (message) greet() {
	fmt.Println("inside the message type")
}

func (m message) print() {
	fmt.Printf("the messsage is = %s\n", m)
}

func (m *message) change(t string) {
	*m = message(t)
	fmt.Println("message was changed")
}

type person struct {
	name string
}

func (person) greet() {
	fmt.Println("hello iam person")
}

func (p person) greetWithName() {
	fmt.Printf("hello iam %s\n", p.name)
}

func (p *person) changeName(n string) {
	p.name = n
	fmt.Println("name was changed")
}

func Methods() {
	var m message = "message one"
	m.greet()
	m.print()
	m.change("iam jossie")
	m.print()

	// save methods
	c := m.change

	c("iam esteban")
	m.print()

	// save functionality of methods
	// methods expressions
	me := (*message).change

	me(&m, "iam bella")
	m.print()

	var msg message = "message two"
	m.greet()
	msg.print()
	me(&msg, "iam fernanda")
	msg.print()
}

func StructsMethods() {
	var p = person{name: "jossie"}
	p.greet()
	p.greetWithName()
	p.changeName("esteban")
	p.greetWithName()
}
