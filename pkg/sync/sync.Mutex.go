package main

import "fmt"
import "sync"
import "time"

func main() {
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var s string
	fmt.Scanln(&s)
}

/*

$ go run sync.Mutex.go
9 start
9 end
0 start
0 end
1 start
1 end
2 start
2 end
3 start
3 end
4 start
4 end
5 start
5 end
6 start
6 end
7 start
7 end
8 start
8 end

*/
