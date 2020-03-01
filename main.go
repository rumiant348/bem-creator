package main

import (
	"fmt"
	"github.com/rumiant348/bem-creator/lib/className"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Script takes 1 argument - classname")
		os.Exit(1)
	}
	input := os.Args[1]
	c := className.NewClassName(input)
	createCss(c)
}

func createCss(c *className.ClassName) {

	currentFolder := getCurrentFolder()
	folderPath := path.Join(currentFolder, c.GetDirPath())
	filePath := path.Join(currentFolder, c.GetFilePath())

	// check if file exists
	f, _ := os.Stat(filePath)
	if f != nil {
		fmt.Printf("%v \nFile exists already", filePath)
		os.Exit(1)
	}

	os.MkdirAll(folderPath, os.ModePerm)
	data := c.GetCssTemplate()
	os.Create(filePath)
	ioutil.WriteFile(filePath, []byte(data), os.ModePerm)
	fmt.Printf("File written for %v\n", filePath)
	fmt.Printf("Import path:\n @import \"%v\";", c.GetClassName())
}

func getCurrentFolder() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
