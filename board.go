package main;

import "fmt"

type Board struct {
    height, width float64
    grid *Grid
    sensors []*Sensor
}


func (board Board) String() string {
    sensor_str := "["
    for _, sensor_ptr := range board.sensors {
        sensor_str = fmt.Sprintf("%s\n%s", sensor_str, *sensor_ptr)
    }
    sensor_str = fmt.Sprintf("%s\n  ]", sensor_str)

    return fmt.Sprintf("Board {\n  h: %f\n  w: %f\n  grid: %s\n  sensors:%s\n}", board.height, board.width, board.grid, sensor_str)
}

