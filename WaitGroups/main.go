package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var printWG sync.WaitGroup
	for i := 0; i < 5; i++ {
		printWG.Add(1)
		go print(i, &printWG)
	}

	//wait tell wg counter 0
	printWG.Wait()
	//time.Sleep(2*time.Second)

}

func print(s int, wg *sync.WaitGroup) {

	//this defer make sure that this function will execute before the end of function
	defer wg.Done()

	fmt.Println("start", s)
	time.Sleep(time.Second)
	fmt.Println("end", s)

}
