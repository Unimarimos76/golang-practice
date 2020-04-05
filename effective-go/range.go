package main

import "fmt"
func main() {

for pos, char := range "日本語" {
  fmt.Printf("charcter %c starts at byte position %d", char, pos)
}
}
