package race_condition

import "testing"

// run with "race detector"
// cd race_condition
// go test -race
func TestReadAndWrite(t *testing.T) {
	ReadAndWrite()
}
