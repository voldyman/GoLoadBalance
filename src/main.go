package main

import (
	"fmt"
	"flag"
)

func main() {
	var backends_input strings

	var port = flag.Int("port", 8050, "The port to listen for connections")

	flag.Var(&backends_input ,"backend", "List of backend servers, need at least one")

	flag.Parse()

	if flag.NFlag() == 0 || len(backends_input) < 1 {
        flag.PrintDefaults()
		return
    } 

	backends := NewBackends()

	for _, addr := range backends_input {
		backends.Add(addr)
	}

	startServer(*port, backends)
}

type strings []string

func (self *strings) String() string {
	return fmt.Sprintf("%s", *self)
}

func (self *strings) Set(val string) error{

	*self = append(*self, val)

	return nil
}
