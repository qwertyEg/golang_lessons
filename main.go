package main

import (
	"fmt"
	"golang_lessons/lesson_4/geometry"
)

func main() {
    p1 := geometry.NewPoint(0, 0)
    p2 := geometry.NewPoint(3, 4)
    fmt.Printf("Distance: %.2f\n", p1.DistanceTo(p2))
} 