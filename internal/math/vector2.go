package math

import (
	"fmt"
	"math"
)

type Vector2 struct {
    X float64
    Y float64
}

func (self *Vector2) Add(other Vector2) {
    self.X += other.X
    self.Y += other.Y
}

func (self *Vector2) AddScalar(other float64) {
    self.X += other
    self.Y += other
}

func (self *Vector2) Sub(other Vector2) {
    self.X -= other.X
    self.Y -= other.Y
}

func (self *Vector2) SubScalar(other float64) {
    self.X -= other
    self.Y -= other
}

func (self *Vector2) Mul(other Vector2) {
    self.X *= other.X
    self.Y *= other.Y
}

func (self *Vector2) MulScalar(other float64) {
    self.X *= other
    self.Y *= other
}

func (self Vector2) Dot(other Vector2) float64 {
    return (self.X * other.X) + (self.Y * other.Y)
}

func (self *Vector2) Div(other Vector2) {
    self.X /= other.X
    self.Y /= other.Y
}

func (self *Vector2) DivScalar(other float64) {
    if other == 0 {
        other = 0.0001
    }

    self.X /= other
    self.Y /= other
}

func (self Vector2) DistanceFrom(other Vector2) float64 {
    return math.Sqrt(
        math.Pow((self.X - other.X), 2) - math.Pow((self.Y - other.Y), 2),
    )
}

func (self Vector2) Len() float64 {
    return math.Sqrt(self.Dot(self))
}

func (self Vector2) Normalized() Vector2 {
    res := self.Copy()
    res.DivScalar(res.Len())

    return res
}

func (self Vector2) Copy() Vector2 {
    return Vector2{
        X: self.X,
        Y: self.Y,
    }
}

func (self *Vector2) String() string {
    return fmt.Sprintf("( %f; %f )", self.X, self.Y)
}
