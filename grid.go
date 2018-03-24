package main;

import "fmt"

type Grid struct {
    n_cells_v int
    n_cells_h int
}

func (grid Grid) String() string {
    return fmt.Sprintf("Grid {\n  n_cells_v: %d\n  n_cells_h: %d\n}", grid.n_cells_v, grid.n_cells_h)
}
