package utils

import (
	"fmt"
	"os"
)

func Loge(a ...interface{}) {
	fmt.Fprintln(os.Stdout, a...)
}
func CheckErr(err error)  {
	panic(err)
}
