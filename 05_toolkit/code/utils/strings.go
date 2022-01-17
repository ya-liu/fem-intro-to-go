package utils

import (
	"fmt"
	"strings"
)

func MakeExcited(str string) string {
	var strUpper = strings.ToUpper(str)
	var output = strUpper + "!"
	fmt.Println(output)
	return output
}
