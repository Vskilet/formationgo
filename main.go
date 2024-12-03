package main

import (
	"fmt"
	"time"
	"warmup/maps"
	"warmup/slices"
	"warmup/types"
	"warmup/users"
)

func main() {
	fmt.Println("My Users")
	user1 := users.Create("SENE", "Victor", time.Now())
	user2 := users.Create("DUGNY", "Lucas", time.Now())
	user3 := users.Create("MAILLOT", "Samuel", time.Now())
	fmt.Println(user1)

	fmt.Println("\nExo slice")
	a := []int{0,1,2,3,4}
	fmt.Printf("This is a : %v\nThis is b : %v\n", a, slices.Copy(a))
	
	fmt.Println("\nExo map")
	ei := maps.New()
	ei.Add(0, user1)
	ei.Add(1, user2)
	ei.Add(2, user3)
	fmt.Println(ei)

	fmt.Println("\nExo types")
	i := types.MyInt(40)
	var n int
	n = 2
	fmt.Printf("i=%v and n=%v\n", i, n)
	fmt.Println(i.Divide(n))
	// fmt.Println("Add : ", i.Add(n))
	// fmt.Println("Sub : ", i.Sub(n))
	// fmt.Println("Multiply : ", i.Multiply(n))
}
