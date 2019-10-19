package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("===== Your ID =====")
	fmt.Println(u.String())
}
