package main;

import (
    "flag"
    "fmt"
    "errors"
)

const DEFAULT_CONFIG_FILE = "/dev/null"

func newFieldGridSensorSimulator() (*FieldGridSensorSimulator, error) {
    var app *FieldGridSensorSimulator = new(FieldGridSensorSimulator)

    flag.IntVar(&(app.run_count), "runs", 1, "number of runs to execute")
    flag.StringVar(&(app.config_file_path), "conf", DEFAULT_CONFIG_FILE, "configuration file path")

    flag.Parse()
    if app.config_file_path == DEFAULT_CONFIG_FILE {
        return nil, errors.New("fatal: no configuration file stated, please use -conf <file> to state a configuration file")
    }
//    app.run_count = *run_count_ptr
    return app, nil
}

type FieldGridSensorSimulator struct {
    run_count int
    config_file_path string
}

func (fgss *FieldGridSensorSimulator) run() int {
    board := newBoard(1, 3, 10, 5)
//    board.populateSensors(0, 10, LinearDistanceFunction2d{4})
    board.populateSensors(0,0, PhysicalDistanceFunction2d{0.35})

    field_pnt := Point2d{1, 2}
    board.generateSensorDataForTarget(field_pnt)
    fmt.Println(board)
    return 0
}
