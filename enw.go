package enw

import (
        "sync"
        "time"
        "os"

        "github.com/joho/godotenv"
)

type Handler func(string, string)

func Watch(name string, handler Handler) {
        lock.Lock()

        flock[name] = &env{
                value: os.Getenv(name),
                handler: handler,
        }
        if len(flock) == 1 {
                start()
        }

        lock.Unlock()
}

func Forget(name string) {
        lock.Lock()

        delete(flock, name)
        if len(flock) == 0 {
                stop()
        }

        lock.Unlock()
}

type (
        env struct {
                value string
                handler Handler
        }
)

var (
        lock sync.RWMutex
        flock map[string]*env = map[string]*env{}

        stopChan chan struct{}
)

func start() {
        stopChan = make(chan struct{})
        go watching()
}

func stop() {
        close(stopChan)
}

func watching() {
        t := time.NewTicker(5*time.Second)
        for {
                select {
                case <-t.C:
                        godotenv.Overload()
                        for n, e := range flock {
                                v := os.Getenv(n)
                                if e.value != v {
                                        lock.Lock()
                                        e.value = v
                                        lock.Unlock()
                                        e.handler(n, e.value)
                                }
                        }
                case <-stopChan:
                        return
                }
        }
}
