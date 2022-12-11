package main

type Monkey struct {
	items      []int
	operation  Operation
	testDiv    int
	throwTrue  int
	throwFalse int

	inspected int
}
