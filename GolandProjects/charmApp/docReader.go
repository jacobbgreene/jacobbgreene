package main

import (
	"fmt"
	"os"
)

func docReader() {
	items, _ := os.ReadDir("charmApp/Poems")
	for _, item := range items {
		fmt.Println(item.Name())
	}
}
