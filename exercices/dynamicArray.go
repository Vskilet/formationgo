package main

import "fmt"

func main(){
	correct()
	moreComplex()
}

func correct(){
	fmt.Println("My answer which is correct")
	a := []string{"A", "B", "C", "D"}
	fmt.Println(a)
	b := make([]string, len(a))
	copy(b, a)
	a[0] = "z"
	fmt.Println(a,b)
}

func moreComplex(){
	fmt.Println("More complex way")
	a := []string{"A", "B", "C", "D"}
	fmt.Println(a)
	
	var b []string
	b = append(b, a...)
	a[0] = "z"
	fmt.Println(a,b)
}