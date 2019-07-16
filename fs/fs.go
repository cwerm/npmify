package fs

import (
	"fmt"
	"npmify/util"
	"os"
)

func CreateDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		err := os.MkdirAll(dirName, 0755)
		util.CheckErr(err)
		return true
	}

	if src.Mode().IsRegular() {
		fmt.Println(dirName, "already exist as a file!")
		return false
	}

	return false
}