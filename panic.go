package main

import "os"

func main() {
	// panic("a problem")

	_, err := os.Create("/a")
	if err != nil {
		panic(err)
	}
}
