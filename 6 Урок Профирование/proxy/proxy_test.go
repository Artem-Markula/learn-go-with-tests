package main

import "testing"

func TestProxy(t *testing.T) {
	var subject ISubject
	subject = new(RealSubject)

	proxy := Proxy{}
	proxy.setSubject(subject)
	proxy.doSomething()
}
