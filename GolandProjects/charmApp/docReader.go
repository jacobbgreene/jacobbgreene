package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func docReader() {
	files, err := os.ReadDir("Poems")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
		exec.Command("glow " + file.Name())
	}

}
