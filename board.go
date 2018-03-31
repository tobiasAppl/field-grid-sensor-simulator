package main;

import (
    "fmt"
    "math"
    "sync"
    "runtime"
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
    last_sensor_values []float64
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
    if board.last_sensor_values != nil || len(board.last_sensor_values) != len(board.sensors) {
        board.last_sensor_values = nil //free allocated space
        runtime.GC() //force garbage collection
    }
    board.last_sensor_values = make([]float64, len(board.sensors))
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

func (board *Board) _generateSensorDataMap(target_pos *Point2d, i0, i1 int, wg *sync.WaitGroup) {
    defer wg.Done()

    if (i1-i0) == 1 {
        board.last_sensor_values[i0] = board.sensors[i0].calculate_field_effect(*target_pos)
    } else {
        var wg_c sync.WaitGroup
        wg_c.Add(2)

        var im int = i0 + ((i1-i0)/2)

        go board._generateSensorDataMap(target_pos, i0, im, &wg_c)
        go board._generateSensorDataMap(target_pos, im, i1, &wg_c)

        wg_c.Wait()
    }
}

func (board *Board) generateSensorDataForTarget(target_pos Point2d) []float64 {
    if board.sensors == nil {
        return board.last_sensor_values
    }
    if board.last_sensor_values != nil || len(board.last_sensor_values) != len(board.sensors) {
        board.last_sensor_values = nil //free allocated space
        runtime.GC() //force garbage collection
    }
    board.last_sensor_values = make([]float64, len(board.sensors))

    var wgT sync.WaitGroup
    wgT.Add(1)

    go board._generateSensorDataMap(&target_pos, 0, len(board.sensors), &wgT)

    wgT.Wait()

    return board.last_sensor_values
}

func (board Board) String() string {
    sensor_str := "["
    for _, sensor_ptr := range board.sensors {
        sensor_str = fmt.Sprintf("%s\n%s", sensor_str, *sensor_ptr)
    }
    sensor_str = fmt.Sprintf("%s\n  ]", sensor_str)

    last_sensor_values_str := "["
    for _, last_sensor_value := range board.last_sensor_values {
        last_sensor_values_str = fmt.Sprintf("%s %f", last_sensor_values_str, last_sensor_value)
    }
    last_sensor_values_str = fmt.Sprintf("%s ]", last_sensor_values_str)

    return fmt.Sprintf("Board {\n  h: %f\n  w: %f\n  grid: %s\n  sensor_arry:%s\n  sensor_data:%s\n}", board.height, board.width, board.grid, sensor_str, last_sensor_values_str)
}

