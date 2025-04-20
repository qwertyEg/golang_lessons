// Package geometry предоставляет базовые операции с геометрическими фигурами.
package geometry

import "math"

// Point представляет точку в 2D-пространстве.
type Point struct {
    X, Y float64
}

// NewPoint создает новую точку с заданными координатами.
func NewPoint(x, y float64) Point {
    return Point{X: x, Y: y}
}

// DistanceTo вычисляет расстояние между текущей точкой и целевой.
func (p Point) DistanceTo(target Point) float64 {
    dx := p.X - target.X
    dy := p.Y - target.Y
    return math.Sqrt(dx*dx + dy*dy)
}

// Circle представляет окружность с центром и радиусом.
type Circle struct {
    Center Point
    Radius float64
}

// Area вычисляет площадь окружности.
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// IsPointInside проверяет, находится ли точка внутри окружности.
func (c Circle) IsPointInside(p Point) bool {
    return c.Center.DistanceTo(p) <= c.Radius
} 