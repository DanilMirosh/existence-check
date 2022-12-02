package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.json")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File not found")
		return
	}

	// выполните какое-нибудь действие с файлом
	defer file.Close()
}
