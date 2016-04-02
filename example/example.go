package main

import (
	_ "github.com/joho/godotenv/autoload"

	"fmt"
	"math/rand"
	"time"

	"github.com/alxmsl/enw"
	"github.com/mgutz/logxi/v1"
)

func main() {
	enw.Watch("LOGXI", func(name string, value string) {
		fmt.Println("LOGXI set to", value)
		log.ProcessLogxiEnv(value)
		log.DefaultLog = log.New("~")
	})

	go process()

	// change .env file here and see how the output will be changed

	//time.Sleep(10*time.Second)
	//enw.Forget("LOGXI")

	// change .env file here and see how the output won't be changed

	fmt.Scanln()
}

func process() {
	for {
		//time.Sleep(100*time.Millisecond)
		time.Sleep(time.Second)
		v := rand.Float32()
		switch {
		case v < 0.1:
			log.Error("error message")
		case v < 0.3:
			log.Info("info message")
		default:
			log.Debug("debug message")
		}
	}
}
