package main;

import "fmt"

func newSensor() *Sensor {
    sensor := new(Sensor)
    sensor.pos = Point2d{0,0}
    sensor.val_min = 0
    sensor.val_max = 1
    sensor.dist_func = LinearDistanceFunc2d{ 1 }
    return sensor
}

type Sensor struct {
    pos Point2d
    val_min, val_max float64
    dist_func DistanceFunc2d
}

func (sensor Sensor) calculate_field_effect(target_pos Point2d) float64  {
    var dist_fac float64 = sensor.dist_func.calcDistanceFactor(sensor.pos, target_pos)
    var val_diff float64 = sensor.val_max - sensor.val_min

    return val_diff * dist_fac
}

func (sensor Sensor) String() string {
    return fmt.Sprintf("Sensor {\n  pos: %s\n val_min: %f\n  val_max: %f\n}", sensor.pos, sensor.val_min, sensor.val_max)
}
