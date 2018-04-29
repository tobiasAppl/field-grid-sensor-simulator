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
//    fmt.Println("new fgss end")
//    app.run_count = *run_count_ptr
    return app, nil
}

type InputPointConfiguration struct {
    points []Point2d
}

type AppConfiguration struct {
    board_height, board_width float64
    nr_cells_vertical, nr_cells_horizontal int
    input_point_file_path, output_file_path string
    randomized bool
}

type FieldGridSensorSimulator struct {
    run_count int
    config_file_path string
    configuration AppConfiguration
    input_points InputPointConfiguration
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

    if len(fgss.configuration.input_point_file_path) > 0 {
        ipfile, ipferr := os.Open(fgss.configuration.input_point_file_path)
        if ipferr != nil {
            return err
        }
        defer ipfile.Close()

        fgss.input_points = InputPointConfiguration{}

        ipfreader := json.NewDecoder(ipfile)
        ipfdecoding_err := ipfreader.Decode(&(fgss.input_points))
        if ipfdecoding_err != nil && ipfdecoding_err != io.EOF {
            return ipfdecoding_err
        }
    }
    return nil
}

func (fgss *FieldGridSensorSimulator) run() (int, error) {
    fmt.Println("run start")
    board := newBoard(fgss.configuration.board_height,
                      fgss.configuration.board_width,
                      fgss.configuration.nr_cells_vertical,
                      fgss.configuration.nr_cells_horizontal)
//    board.populateSensors(0, 10, LinearDistanceFunction2d{4})
    board.populateSensors(0,0, PhysicalDistanceFunction2d{0.35})
    fmt.Println("run start")

    field_pnt := Point2d{1, 2}
    update_err := board.updateSensorDataForTarget(field_pnt)
    if update_err != nil {
        return 1, update_err
    }
    fmt.Printf("{\n \"board\": %s\n", board)

    fmt.Printf("config: %s", fgss.configuration)
    return 0, nil
}
