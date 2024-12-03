package main

import (
	"fmt"
	"sort"
)

func main(){
	firstMap()
}

func firstMap(){
	a := map[int]int{0:1, 1:2, 2:3}
	for i := 0; i < len(a); i++ {
		fmt.Println(a)
	}
}

func order(){
	var m map[int]string{0:"a", 1:"b", 2:"c"}
	var k []int
}