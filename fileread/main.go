package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("start")
	f, err := os.Open("test")
	if err != nil {
		fmt.Println("error")
	}

	defer f.Close()

	buf := make([]byte, 5000)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
