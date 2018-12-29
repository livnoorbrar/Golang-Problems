package main

import "fmt"

func main() {
	s := []int{0,1,2,3,4,5,6,7,8}
	fmt.Println(rotateLeft(s,8))
}

func rotateLeft(slice []int, places int) []int {
	ns := make([]int,len(slice))
	latterPart := slice[places:]
	copy(ns,latterPart)
	copy(ns[len(latterPart):],slice[:places])
	return ns
}
