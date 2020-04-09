package main

import "fmt"

type ISubject interface {
	doSomething()
}

type RealSubject struct {
}

func (*RealSubject) doSomething() {
	fmt.Println("realsubject do something here!")
}

type Proxy struct {
	realSubject ISubject
}

func (p *Proxy) setSubject(s ISubject) {
	p.realSubject = s
}

func (p *Proxy) doSomething() {
	
	fmt.Println("proxy do something here!")
	p.realSubject.doSomething()
}
