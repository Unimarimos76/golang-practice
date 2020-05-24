package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("/Users/uni880/golang-practice/stat/hoge.txt")
	if err != nil {
		log.Fatal(err)
	}
}
