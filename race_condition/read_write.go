package race_condition

func ReadAndWrite() {
	var value int
	var count int
	go func() {
		for {
			count += value
		}
	}()

	for i := 0; i < 10; i += 1 {
		value = i
	}
}
