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

func WriteToFile(file, content string) error {
	fo, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer fo.Close()
	_, err = fo.WriteString(content)
	return err
}
