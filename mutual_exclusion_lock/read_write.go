package mutual_exclusion_lock

import "sync"

func ReadAndWrite() {
	var mux sync.Mutex

	var value int
	var count int
	go func() {
		for {
			mux.Lock()
			count += value // read
			mux.Unlock()
		}
	}()

	for i := 0; i < 10; i += 1 {
		mux.Lock()
		value = i // write
		mux.Unlock()
	}
}
