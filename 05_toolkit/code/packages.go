package main

import (
	"fem-intro-to-go/05_toolkit/code/utils"
	"fmt"
)

func calculateData() int {
	totalValue := utils.Add(2, 3, 4)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateData()
	fmt.Println(total)
}

// at first, 'go run packages.go' yields this error: package fem-intro-to-go/05_toolkit/code/utils is not in GOROOT
// researched online and used command 'go env -w GO111MODULE=off'
// read more at https://maelvls.dev/go111module-everywhere/
