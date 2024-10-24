package Generator

func FirstGeneratorId() func() int {
	id := 1000
	return func() int {
		id++
		return id
	}
}

func SecondGeneratorId() func() int {
	id := 10000
	return func() int {
		id++
		return id
	}
}
