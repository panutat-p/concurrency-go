package race_condition

import "testing"

// run with 'race detector'
// go test ./race_condition -race
func TestReadAndWrite(t *testing.T) {
	ReadAndWrite()
}
