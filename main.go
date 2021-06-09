package main

import (
	"fmt"
	"time"
)

func PrintTime() string {
	return fmt.Sprintf("Current Time Is %s: ", time.Now())
}

func main() {
	fmt.Printf(PrintTime())

	str := fmt.Sprintf("Hey There! I am intentionally failing the build by not using variable str")
}
