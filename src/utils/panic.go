package utils

import "fmt"

func PanicIfNotNil(err error, format string) {
	if err != nil {
		panic(fmt.Errorf(format, err))
	}
}

func Panic(s string) {
	panic(fmt.Errorf(s))
}
