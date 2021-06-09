package main

import (
	"fmt"
	"time"
)

func PrintTime() string {
           	return fmt.Sprintf("Current Time Is %s: ", time.Now())
}

func main() {
	       fmt.Println(PrintTime())

}
