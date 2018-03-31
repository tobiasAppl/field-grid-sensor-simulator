package main;

import "fmt"

func main() {

    board := newBoard(1, 3, 10, 5)
//    board.populateSensors(0, 10, LinearDistanceFunction2d{4})
    board.populateSensors(0,0, PhysicalDistanceFunction2d{0.35})

    field_pnt := Point2d{1, 2}
    board.generateSensorDataForTarget(field_pnt)
    fmt.Println(board)
}

