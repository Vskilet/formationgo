package main

import (
	"fmt"
	"warmup/users"
	"time"
)

func main() {
	user1 := users.Create("SENE", "Victor", time.Now())
	fmt.Println(user1)
}
