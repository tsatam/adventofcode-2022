package main

func processRounds(monkeys []Monkey, rounds int, worry bool) []int {
	modFactor := 1
	for _, monkey := range monkeys {
		modFactor *= monkey.testDiv
	}

	for i := 0; i < rounds; i++ {
		for monkeyIdx, monkey := range monkeys {
			for ; len(monkey.items) > 0; monkey.items = monkey.items[1:] {
				item := monkey.items[0]
				item = monkey.operation.process(item)

				if !worry {
					item /= 3
				}

				item = item % modFactor

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

	inspected := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspected[i] = monkey.inspected
	}

	return inspected
}
