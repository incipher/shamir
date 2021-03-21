package main

import (
	"fmt"

	"github.com/hashicorp/vault/shamir"
)

func main() {
	shards, err := shamir.Split([]byte("hello"), 3, 2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(shards)
}
