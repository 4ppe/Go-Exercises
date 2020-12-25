package main

import (
	"fmt"

	"github.com/4ppe/Go-Exercises/slices/limit-the-backing-array-sharing/solution/api"
)

func main() {

	received := api.Read(0, 3)

	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.received :", received)
}
