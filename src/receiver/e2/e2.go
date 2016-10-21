package main

import "log"

type customError struct{}

func (customError) Error() string {
	return ""
}

func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error = nil
	c, err := fail()
	if err != nil {
		log.Fatal("FAIL")
	}

	println("OK")
	_ = c
}
