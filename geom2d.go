/**
 * Author: Tobias Appl <tobias.appl@gmail.com>
**/

package main;

import "fmt"
import "math"

type Point2d struct {
    x, y float64
}

func (p2d Point2d) String() string {
    return fmt.Sprintf("Point2d {\n  x:%f\n  y:%f\n}", p2d.x, p2d.y)
}

func (p2d Point2d) substract(other Point2d) Point2d {
    return Point2d{p2d.x - other.x, p2d.y - other.y}
}

func (p2d Point2d) add(other Point2d) Point2d {
    return Point2d{p2d.x + other.x, p2d.y + other.y}
}

func (p2d Point2d) length() float64 {
    return  math.Sqrt(math.Pow(p2d.x, 2) + math.Pow(p2d.y, 2))
}
