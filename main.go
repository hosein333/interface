package main

import (
	"fmt"
	"time"
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
func printLog(pChannel chan interface{}) {
    for {
        v, ok := <-pChannel

        if !ok {
            fmt.Println("pChannel Closed")
            return
        }

        fmt.Println(v)
    }
}
func main() {
    s := Square{Length: 10}
    r := Rectangle{Lenght: 10, Width: 20}

        LogChannel := make(chan interface{}, 10)

            go printLog(LogChannel)

            time.Sleep(1 * time.Second)


            LogChannel <- "test"
            LogChannel <- "1234"
            LogChannel <- "!@#4"
            LogChannel <- "t1!"

    
    shapeList := []Shape{s, r}

    for _, value := range shapeList {
        fmt.Printf("\ntype : %T, Area : %d\n", value, value.Area())
        square, ok := value.(Square)
        fmt.Printf("type: %T , OK : %v\n", square, ok)
    }
    res := CalculateArea(s)
    fmt.Println(res)
}

// CalculateArea godoc
func CalculateArea(s Shape) int {
    return s.Area()
}
