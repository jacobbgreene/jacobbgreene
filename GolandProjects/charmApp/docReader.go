package charmApp

import (
	"fmt"
	"os"
)

func docReader() {
	items, _ := os.ReadDir(".")
	for _, item := range items {
		if item.IsDir() {
			subitems, _ := os.ReadDir(item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					// handle file there
					fmt.Println(item.Name() + "/" + subitem.Name())
				}
			}
		} else {
			// handle file there
			fmt.Println(item.Name())
		}
	}
}
