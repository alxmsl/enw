package main

import (
        "fmt"
        "time"

        _ "github.com/joho/godotenv/autoload"
        "github.com/mgutz/logxi/v1"

        "github.com/alxmsl/enw"
)

func main() {
        enw.Watch("LOGXI", func(name string, value string) {
                fmt.Println("LOGXI set to", value)
                log.ProcessLogxiEnv(value)
                log.DefaultLog = log.New("~")
        })

        go process()
        fmt.Scanln()
}

func process() {
        for {
                log.Info("info message")
                log.Debug("debug message")
                time.Sleep(time.Second)
        }
}
