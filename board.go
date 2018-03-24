package main;

import (
    "fmt"
    "math"
)

/*func newBoard() *Board {
    board := new(Board)
    board.height = 1
    board.width = 1
    grid := new(Grid)
    grid.n_cells_v = 1
    grid.n_cells_h = 1
    board.grid = grid
    return board
}*/

func newBoard(height, width float64, n_cells_v, n_cells_h int) *Board {
    board := new(Board)
    board.height = height
    board.width = width
    grid := new(Grid)
    grid.n_cells_v = n_cells_v
    grid.n_cells_h = n_cells_h
    board.grid = grid
    return board
}

type Board struct {
    height, width float64
    grid *Grid
    sensors []*Sensor
}

func (board *Board) getCellNrForTargetPosition(target_pos Point2d) int {
    if board.height <= 0 || board.width == 0 ||
        target_pos.x < 0 || board.width <= target_pos.x ||
        target_pos.y < 0 || board.height <= target_pos.y ||
        board.grid == nil {
        return 0
    }
    var u int = int(math.Floor((target_pos.x * float64(board.grid.n_cells_h))/board.width))
    var v int = int(math.Floor((target_pos.y * float64(board.grid.n_cells_v))/board.height))

    var n_cell int = v * board.grid.n_cells_h + u + 1
    return n_cell
}

func (board *Board) populateSensors(val_min, val_max float64, dist_func DistanceFunc2d) {
    if board.grid == nil {
        return
    }
    var n_sens_v int = board.grid.n_cells_v + 1
    var n_sens_h int = board.grid.n_cells_h + 1

    var cell_height float64 = board.height / float64(board.grid.n_cells_v)
    var cell_width float64 = board.width / float64(board.grid.n_cells_h)

    for i_v := 0; i_v < n_sens_v; i_v++ {
        for i_h := 0; i_h < n_sens_h; i_h++ {
            var sensor *Sensor = new(Sensor)
            sensor.val_min = val_min
            sensor.val_max = val_max
            sensor.dist_func = dist_func

            var pos Point2d
            pos.x = float64(i_v) * cell_width
            pos.y = float64(i_h) * cell_height

            sensor.pos = pos

            board.sensors = append(board.sensors, sensor)
        }
    }
}

func (board *Board) setSensorMinValue(val_min float64) {
    if board.sensors == nil {
        return
    }
    for _, v := range board.sensors {
        v.val_min = val_min
    }
}

func (board *Board) setSensorMaxValue(val_max float64) {
    if board.sensors == nil {
        return
    }
    for _, v := range board.sensors {
        v.val_max = val_max
    }
}

func (board *Board) setSensorEffectDistanceFunction(dist_func DistanceFunc2d) {
    if board.sensors == nil {
        return
    }
    for _, v := range board.sensors {
        v.dist_func = dist_func
    }
}

func (board *Board) generateSensorDataForTarget(target_pos Point2d) []float64 {
    var sensor_values []float64
    if board.sensors == nil {
        return sensor_values
    }

    for _, sensor := range board.sensors {
        var value float64 = sensor.calculate_field_effect(target_pos)
        sensor_values = append(sensor_values, value)
    }

    return sensor_values
}

func (board Board) String() string {
    sensor_str := "["
    for _, sensor_ptr := range board.sensors {
        sensor_str = fmt.Sprintf("%s\n%s", sensor_str, *sensor_ptr)
    }
    sensor_str = fmt.Sprintf("%s\n  ]", sensor_str)

    return fmt.Sprintf("Board {\n  h: %f\n  w: %f\n  grid: %s\n  sensors:%s\n}", board.height, board.width, board.grid, sensor_str)
}

