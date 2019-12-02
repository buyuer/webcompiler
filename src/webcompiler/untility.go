package webcompiler

import (
	"fmt"
	"os"
)

func checkError(err error) error {
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func isFileExists(path string) (result bool) {
	_, err := os.Stat(path)
	result = err == nil || os.IsExist(err)
	return
}
