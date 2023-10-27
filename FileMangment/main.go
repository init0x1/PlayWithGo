package main

import (
	"fmt"
	"init0x1/FileMangment/fileutils"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	//use readtextfile
	content, err := fileutils.ReadTextFile(rootPath + "/Files/flag.txt")
	//use WriteToFile func
	newContent := "test from write to file function "
	fileutils.WriteToFile(rootPath+"/Files/output.txt", newContent)
	if err == nil {
		fmt.Printf(content)
	} else {
		fmt.Printf("Error While Reading File %v", err)
	}

}
