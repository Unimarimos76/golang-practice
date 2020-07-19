package main

import "io/ioutil"

func main() {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
}
