package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()
			fmt.Println(count)
			count++
			m.Unlock()
		}()
	}
	var input string
	fmt.Scanln(&input)
}
