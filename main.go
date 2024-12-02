package main

import (
	"fmt"
	"warmup/user"
	"time"
)

func main() {
	user1 := user.Create("SENE", "Victor", time.Now())
	fmt.Println(user1)
}
