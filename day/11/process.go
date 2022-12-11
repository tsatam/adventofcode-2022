package main

func processRounds(monkeys []Monkey, rounds int) []Monkey {
	for i := 0; i < rounds; i++ {
		for monkeyIdx, monkey := range monkeys {
			for ; len(monkey.items) > 0; monkey.items = monkey.items[1:] {
				item := monkey.items[0]
				item = monkey.operation.process(item)
				item /= 3

				if item%monkey.testDiv == 0 {
					monkeys[monkey.throwTrue].items = append(monkeys[monkey.throwTrue].items, item)
				} else {
					monkeys[monkey.throwFalse].items = append(monkeys[monkey.throwFalse].items, item)
				}
				monkey.inspected = monkey.inspected + 1
			}
			monkeys[monkeyIdx] = monkey
		}
	}
	return monkeys
}
