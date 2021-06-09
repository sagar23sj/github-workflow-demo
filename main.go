package main

import (
	"fmt"
	"time"
)

func PrintTime() string {
	return fmt.Sprintf("Current Time Is %s: ", time.Now())
}

func processValue(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			processValue(i)
		}()
	}
	fmt.Println(PrintTime())

}
