package main;

import (
    "flag"
    "fmt"
    "errors"
    "encoding/json"
    "os"
    "io"
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

    parse_err := app.parse_config_json()
    if parse_err != nil {
        return nil, parse_err
    }
    fmt.Println("new fgss end")
//    app.run_count = *run_count_ptr
    return app, nil
}

type AppConfiguration struct {

}

type FieldGridSensorSimulator struct {
    run_count int
    config_file_path string
    configuration AppConfiguration
}

func (fgss *FieldGridSensorSimulator) parse_config_json() error {
    cfile, err := os.Open(fgss.config_file_path)
    if err != nil {
        return err
    }
    defer cfile.Close()
    fgss.configuration = AppConfiguration{}

    json_reader := json.NewDecoder(cfile)
    decoding_err := json_reader.Decode(&(fgss.configuration))
    if decoding_err != nil && decoding_err != io.EOF {
        return decoding_err
    }

    return nil
}

func (fgss *FieldGridSensorSimulator) run() (int, error) {
    fmt.Println("run start")
    board := newBoard(1, 3, 10, 5)
//    board.populateSensors(0, 10, LinearDistanceFunction2d{4})
    board.populateSensors(0,0, PhysicalDistanceFunction2d{0.35})
    fmt.Println("run start")

    field_pnt := Point2d{1, 2}
    update_err := board.updateSensorDataForTarget(field_pnt)
    if update_err != nil {
        return 1, update_err
    }
    fmt.Printf("board: %s\n", board)
    return 0, nil
}
