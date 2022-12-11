package main

type Operation interface {
	process(old int) int
}

type AddOperation struct {
	value int
}
type MulOperation struct {
	value int
}

type PowOperation struct{}

func (o AddOperation) process(old int) int {
	return old + o.value
}

func (o MulOperation) process(old int) int {
	return old * o.value
}

func (o PowOperation) process(old int) int {
	return old * old
}
