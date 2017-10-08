package main

import "os"

func oink() {
	f, err := os.Open("foo")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close() // DEFER
	defer func() {
		os.Open("foo") // UNCHECKED
	}()
}
