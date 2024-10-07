package main

import (
	"flag"
	"fmt"

	"github.com/lucaslimafernandes/go-ci-actions/utils"
)

func main() {

	nb := flag.Int("n", 0, "Some number to show tabuada")
	flag.Parse()

	fmt.Println("Tabuada:")

	for i := 0; i <= 10; i++ {
		_value := utils.ValuesGenerator(*nb, i)
		fmt.Printf("%v X %v = %v\n", *nb, i, _value)
	}

}
