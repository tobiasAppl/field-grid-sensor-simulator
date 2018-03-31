package main;

import "math"

type PhysicalDistanceFunction2d struct {
    lambda float64
}


func (pdf PhysicalDistanceFunction2d) calcDistanceFactor(sensor *Sensor, target_pos Point2d) float64 {
    if pdf.lambda == 0 {
        return math.MaxFloat64
    }

    var diffVec Point2d = sensor.pos.substract(target_pos)
    var diffVecLen float64 = diffVec.length()

    var incpower = ( 20 * math.Log((4*diffVecLen*math.Pi)/pdf.lambda) ) / 22.6

    return incpower
}
