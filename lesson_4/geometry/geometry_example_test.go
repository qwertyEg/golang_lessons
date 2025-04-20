package geometry_test

import (
	"fmt"
	"golang_lessons/lesson_4/geometry"
)

func ExampleCircle_Area() {
    c := geometry.Circle{
        Center: geometry.NewPoint(0, 0),
        Radius: 5,
    }
    fmt.Printf("Площадь окружности: %.2f", c.Area())
    // Output: Площадь окружности: 78.54
}

func ExamplePoint_DistanceTo() {
    p1 := geometry.NewPoint(0, 0)
    p2 := geometry.NewPoint(3, 4)
    fmt.Printf("Расстояние между точками: %.2f", p1.DistanceTo(p2))
    // Output: Расстояние между точками: 5.00
}

func ExampleCircle_IsPointInside() {
    c := geometry.Circle{
        Center: geometry.NewPoint(0, 0),
        Radius: 5,
    }
    p := geometry.NewPoint(3, 4)
    fmt.Printf("Точка внутри окружности: %v", c.IsPointInside(p))
    // Output: Точка внутри окружности: true
} 