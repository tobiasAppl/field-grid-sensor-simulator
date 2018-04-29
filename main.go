package main;

import (
    "os"
    "log"
    "fmt"
)

func main() {
    errlog := log.New(os.Stderr, "", 0)
    app, err := newFieldGridSensorSimulator()
    if err != nil {
        errlog.Println(err)
        os.Exit(1)
    }

    retval, err := app.run()
    if err != nil {
        errlog.Println(err)
    }
    fmt.Println("pre exit")
    os.Exit(retval)
}

