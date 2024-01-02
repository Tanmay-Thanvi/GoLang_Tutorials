// geometry/circle.go

package geometry

// Circle struct representing a circle
type Circle struct {
    Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
    return Pi * c.Radius * c.Radius
}