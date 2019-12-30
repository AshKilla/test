package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hello(){
	fmt.Println("hello 娜扎")
	wg.Done()
}


func main(){
	wg.Add(1)
	go Hello()
	fmt.Println("hello main")
	wg.Wait()
}
