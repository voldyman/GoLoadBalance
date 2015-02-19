package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	var backends_input strings

	var port = flag.Int("port", 8050, "The port to listen for connections")
	var debug = flag.Bool("debug", false, "Enable debug mode")

	flag.Var(&backends_input, "backend", "List of backend servers, need at least one")

	flag.Parse()

	if flag.NFlag() == 0 || len(backends_input) < 1 {
		flag.PrintDefaults()
		return
	}

	backends := NewBackends()

	backends.Add([]string(backends_input)...)

	if *debug {
		go debugRoutine()
	}

	startServer(*port, backends)
}

func debugRoutine() {
	for {
		<-time.After(2 * time.Second)
		fmt.Println(time.Now(), "NumGoroutine", runtime.NumGoroutine())
	}
}

type strings []string

func (self *strings) String() string {
	return fmt.Sprintf("%s", *self)
}

func (self *strings) Set(val string) error {

	*self = append(*self, val)

	return nil
}
