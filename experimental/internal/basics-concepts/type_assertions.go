package internal

import "fmt"

type canines interface {
	introduceYourself()
}

type dog struct{}

func (d dog) introduceYourself() {
	fmt.Println("iam a dog")
}

func (d dog) bark() {
	fmt.Println("wau wau wau")
}

type wolf struct{}

func (w wolf) introduceYourself() {
	fmt.Println("iam a wolf")
}

func (w wolf) growl() {
	fmt.Println("grr grr grr")
}

func TypeAssertions() {
	var d = dog{}
	var w = wolf{}

	areYouADog(d)
	areYouADog(w)

	speakMe(d)
	speakMe(w)
}

func areYouADog(c canines) {
	c.introduceYourself()

	if d, ok := c.(dog); ok {
		d.bark()
	}
}

func speakMe(c canines) {
	c.introduceYourself()

	switch c := c.(type) {
	case dog:
		c.bark()
	case wolf:
		c.growl()
	default:
		fmt.Println("iam don't known which canine is")
	}
}
