// geometry/rectangle.go

package geometry

// Rectangle struct representing a rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}