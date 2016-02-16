package main

import (
	"fmt"

	"github.com/dmitris/deptesthuge/deptestsmall"
)

func main() {
	s := deptestsmall.SillyType{}.String()
	fmt.Printf("[INFO]: %s\n", s)
}
