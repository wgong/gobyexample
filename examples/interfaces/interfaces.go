// _Interfaces_ are named collections of method
// signatures.

package main

import "fmt"
import "math"

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

type triangle struct {
	a, b, c float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// implementation fo triangle
func (t triangle) perim() float64 {
	return t.a + t.b + t.c
}

func (t triangle) area() float64 {
	// Heron's formula
	semi := 0.5 * (t.a + t.b + t.c)
	return math.Sqrt(semi * (semi - t.a) * (semi - t.b) * (semi - t.c))
}

// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
// behavior of abstract class
type Animal interface {
	Speak() string
}

// implemented by each object sharing the same behavior
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
	fmt.Printf("geometry type: %T\n", g)
	fmt.Println("geometry param: ", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perimeter: ", g.perim())
	fmt.Println("===============")
}

func main() {
	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.
	r := rect{width: 3, height: 4}
	measure(r)

	c := circle{radius: 5}
	measure(c)

	// a triangle
	t := triangle{a: 5, b: 12, c: 13}
	measure(t)

	// use Animal interface
	animals := []Animal{Dog{}, Cat{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Printf("%T \t speaks %s\n", animal, animal.Speak())
	}

}
