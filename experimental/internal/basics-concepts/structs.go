package internal

import "fmt"

type point struct {
	x, y float64
}

type circle struct {
	point
	radio float64
}

type wheel struct {
	circle
	rayos int8
}

func Structs() {
	var p = point{
		x: 2,
		y: 5,
	}

	var c = circle{
		point: p,
		radio: 55,
	}

	var w = wheel{
		circle: c,
		rayos:  55,
	}

	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", w.x)

	pointer := p
	number := 55
	ref := &number

	fmt.Println(pointer, &pointer, *&pointer)
	fmt.Println(number, &number)
	fmt.Println(ref, *ref)
}
