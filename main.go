package main

import (
	"fmt"
	"time"
)

func main() {

	prints := []string{"this is just a git test", "the discreption is not that valid", "don't mind me just continue"}

	for _, pr := range prints {
		fmt.Println(pr)
		time.Sleep(time.Second)
	}
}
