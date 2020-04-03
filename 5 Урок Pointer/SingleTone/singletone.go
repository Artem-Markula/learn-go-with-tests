package main

import (
	"sync"
)



type SingleTone struct {
	instance *SingleTone
	once     sync.Once
}

func (p *SingleTone) GetInstance() *SingleTone{
	p.once.Do(func() {
		p.instance = &SingleTone{}
	})
	return p.instance
}
