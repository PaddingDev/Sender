package utils

import (
	"io/ioutil"
	"os"
)

func ReadFileToByte(file string) (bs []byte, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
