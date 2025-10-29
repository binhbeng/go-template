package main

import (
	"fmt"

	"github.com/binhbeng/goex/cmd"
	"github.com/binhbeng/goex/data"
)

func main() {
	fmt.Println("Hello, World!")
	cmd.Execute()
	data.InitData()
}