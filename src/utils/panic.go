package utils

import "fmt"

func PanicIfNotNil(err error, format string) {
	if err != nil {
		panic(fmt.Errorf(format, err))
	}
}
