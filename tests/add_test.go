package main

import "testing"

func add(x int, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	result := add(1, 2)

	if result != 3 {
		t.Errorf("Expected add(1,2) to be 3, got %d instead", result)
	}
}

/* https://stackoverflow.com/questions/28240489/golang-testing-no-test-files */
/* https://maex.me/2019/03/go-a-comprehensive-introduction/ */
/*

debug@lab:~/Projects/hello-go/tests$ go test 
PASS
ok      _/home/debug/Projects/hello-go/tests    0.001s

*/