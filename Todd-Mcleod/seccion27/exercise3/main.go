package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup


func main() {
	
	contador := 0

	grs := 100
	wg.Add(grs)
	for i := 0; i< grs; i++ {
		go func(){
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(contador)

}
