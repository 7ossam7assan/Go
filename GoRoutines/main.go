package main

import (
	"fmt"
	"time"
)

func main() {
	//print("First")
	//print("Second")
	//go print("First")
	//go print("Second")
	//
	//go func(s string) {
	//	for i:=0;i<5;i++ {
	//		time.Sleep(200* time.Millisecond)
	//		fmt.Println(s,i)
	//	}
	//}("third")

	//testing go in for loop

	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)

}

func print(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s, i)
	}
}
