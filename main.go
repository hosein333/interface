package main

import (
	"fmt"
)

//Shape shape interface
type Shape interface {
    Area() int
}

//Square class
type Square struct {
    Length int
}

//Area implementation for square class
func (s Square) Area() int {
    return s.Length * s.Length
}
// Rectangle class
type Rectangle struct {
    Lenght int
    Width  int
}
// Area imlementation for square class
func (r Rectangle) Area() int {
    return r.Lenght * r.Width
}

func main() {
    s := Square{Length: 10}
    r := Rectangle{Lenght: 10, Width: 20}
    
    shapeList := []Shape{s, r}

    for _, value := range shapeList {
        fmt.Printf("type : %T, Area : %d\n", value, value.Area())
    }
}
