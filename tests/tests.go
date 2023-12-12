package main

import (
	"github.com/justinbather/osprey"
)

func main() {
	o := osprey.New("123")
	if 5 != 10 {

		o.Log("error message test, default")

		o.Critical("critical error test")
	}
}
