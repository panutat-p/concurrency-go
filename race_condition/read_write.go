package race_condition

// ReadAndWrite
// cd race_condition
// go test -race
// build with the so-called "race detector" enabled.
// Data races are a programming error and any program with a data-race is invalid and its behaviour is undefined.

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
