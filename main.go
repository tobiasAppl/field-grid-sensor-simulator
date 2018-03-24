package main;

import "fmt"

func main() {

    board := newBoard(3, 3, 3, 3)
    board.populateSensors(0, 1, LinearDistanceFunc2d{1})
    fmt.Println(board)
}

