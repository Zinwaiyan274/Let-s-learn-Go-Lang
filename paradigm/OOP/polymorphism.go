package OOP

import "fmt"

// Shape is an interface representing a geometric shape
type Shape interface {
	area() float64
}

// Rectangle is a struct representing a rectangle
type Rectangle struct {
	width  float64
	height float64
}

// Circle is a struct representing a circle
type Circle struct {
	radius float64
}

// Implementing the area method for Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Implementing the area method for Circle
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func polymorphism() {
	// Creating instances of Rectangle and Circle
	r := Rectangle{width: 3, height: 4}
	c := Circle{radius: 5}

	// Storing instances in a slice of Shape
	shapes := []Shape{r, c}

	// Iterating through the shapes and calling the area method
	for _, shape := range shapes {
		fmt.Println("Area:", shape.area())
	}
}
