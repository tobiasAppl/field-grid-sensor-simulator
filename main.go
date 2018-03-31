package main;

import "fmt"

func main() {

    board := newBoard(3, 3, 3, 3)
    board.populateSensors(0, 10, LinearDistanceFunc2d{4})

    field_pnt := Point2d{1, 2}
    board.generateSensorDataForTarget(field_pnt)
    fmt.Println(board)
}

