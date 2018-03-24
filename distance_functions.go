package main;

import "fmt"
import "math"

type DistanceFunc2d interface {
    calcDistanceFactor(p0, p1 Point2d) float64
}

type LinearDistanceFunc2d struct {
    max_distance float64
}

func (ldf LinearDistanceFunc2d) calcDistanceFactor(p0, p1 Point2d) float64 {
    if ldf.max_distance <= 0 {
        return 0
    }
    var diffVec Point2d = p0.substract(p1)
    var diffVecLen float64 = diffVec.length()

    var lenNorm = 1 - math.Abs(diffVecLen/ldf.max_distance)

    if lenNorm <= 0 {
        return 0
    }
    return lenNorm
}

func (ldf LinearDistanceFunc2d) String() string {
    return fmt.Sprintf("LinearDistanceFunc2d {\n  max_distance: %f\n}", ldf.max_distance)
}
