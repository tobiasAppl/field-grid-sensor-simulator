package main;

import "fmt"
import "math"

type DistanceFunction2d interface {
    calcDistanceFactor(sensor *Sensor, target_pos Point2d) float64
}

type LinearDistanceFunction2d struct {
    max_distance float64
}

func (ldf LinearDistanceFunction2d) calcDistanceFactor(sensor *Sensor, target_pos Point2d) float64 {
    if ldf.max_distance <= 0 {
        return 0
    }
    var diffVec Point2d = sensor.pos.substract(target_pos)
    var diffVecLen float64 = diffVec.length()
    var lenNorm float64 = 1 - math.Abs(diffVecLen/ldf.max_distance)

    if lenNorm <= 0 {
        return 0
    }
    var val_diff float64 = sensor.val_max - sensor.val_min

    return sensor.val_min + (val_diff * lenNorm)
}

func (ldf LinearDistanceFunction2d) String() string {
    return fmt.Sprintf("LinearDistanceFunc2d {\n  max_distance: %f\n}", ldf.max_distance)
}

