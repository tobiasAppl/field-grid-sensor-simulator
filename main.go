package main;

import (
    "os"
    "fmt"
)

func main() {
    board := newBoard(20,
                      9,
                      100,
                      45)
    board.populateSensors(0,10, PhysicalDistanceFunction2d{0.35})

    field_pnt := Point2d{5, 7}
    update_err := board.updateSensorDataForTarget(field_pnt)
    if update_err != nil {
        fmt.Printf("Error: %s", update_err)
        os.Exit(1)
    }
    fmt.Printf("{\n \"board\": %s\n", board)

    os.Exit(0)
}

