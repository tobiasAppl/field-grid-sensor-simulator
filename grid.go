package main;

import "fmt"

type Grid struct {
    subd_vertical int
    subd_horizontal int
}

func (grid Grid) String() string {
    return fmt.Sprintf("Grid {\n  subd_vert: %d\n  subd_hor: %d\n}", grid.subd_vertical, grid.subd_horizontal)
}
