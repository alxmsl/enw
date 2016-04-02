package enw

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gopkg.in/fsnotify.v1"
)

type (
	Handler func(string, string)
)

type env struct {
	value   string
	handler Handler
}

type cache struct {
	sync.Mutex
	data map[string]*env
}

var (
	vs cache = cache{
		data: map[string]*env{},
	}
	stopChan chan struct{}
)

func Watch(name string, handler Handler) {
	vs.Lock()
	vs.data[name] = &env{
		value:   os.Getenv(name),
		handler: handler,
	}
	if len(vs.data) == 1 {
		start()
	}
	vs.Unlock()
}

func Forget(name string) {
	vs.Lock()
	delete(vs.data, name)
	if len(vs.data) == 0 {
		stop()
	}
	vs.Unlock()
}

func start() {
	stopChan = make(chan struct{})
	go watching()
}

func stop() {
	close(stopChan)
}

func watching() {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic("unable to create new watcher")
	}
	w.Add(".env")

	for {
		select {
		case event := <-w.Events:
			fmt.Println("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("modified file:", event.Name)
				process(event.Name)
			}
		case err := <-w.Errors:
			fmt.Println("error:", err)
		case <-stopChan:
			err := w.Close()
			if err != nil {
				panic("unable to close watcher")
			}
			return
		}
	}
}

func process(filename string) {
	godotenv.Overload(filename)
	vs.Lock()
	for n, e := range vs.data {
		v := os.Getenv(n)
		if e.value != v {
			e.value = v
			e.handler(n, e.value)
		}
	}
	vs.Unlock()
}
