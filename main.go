package main;

import "os"
//import "fmt"
import "log"

func main() {
    errlog := log.New(os.Stderr, "", 0)
    app, err := newFieldGridSensorSimulator()
    if err != nil {
        errlog.Println(err)
        os.Exit(1)
    }

    os.Exit(app.run())
}

