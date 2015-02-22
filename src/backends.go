package main

import (
	"sync"
)

type Backends struct {
	Length    int
	current   int
	Addresses []string

	sync.Mutex
}

func NewBackends() *Backends {
	addresses := []string{}
	current := 0
	length := 0

	return &Backends{
		Length:    length,
		current:   current,
		Addresses: addresses,
	}
}

func (self *Backends) NextAddress() (addess string) {
	self.Lock()

	index := self.current

	self.current = self.current + 1
	if self.current > self.Length-1 {
		self.current = 0
	}

	self.Unlock()
	return self.Addresses[index]
}

func (self *Backends) Add(addresses ...string) {
	self.Lock()

	for _, item := range addresses {
		self.Addresses = append(self.Addresses, item)
	}
	self.Length = len(self.Addresses)

	self.Unlock()
}
