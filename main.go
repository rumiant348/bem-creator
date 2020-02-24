package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//var testWorkfolder string = "./home/rum348/node/YandexPraktikum_task2/blocks"
var workfolder string = "/home/rum348/node/YandexPraktikum_task2/blocks"

func main() {
	if len(os.Args) < 2 {
		panic("Provide argument")
	}
	input := os.Args[1]
	err := createCss(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File written for %v\n", input)
}

func createCss(s string) error {
	a := strings.Split(s, "__")
	class, element := a[0], a[1]
	folderPath := path.Join(workfolder, class, "__" + element)
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return err
	}
	filePath := path.Join(folderPath, fmt.Sprintf("%v__%v.css", class, element))

	// check that file exists
	f, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if f != nil {
		return os.ErrExist
	}
	data := fmt.Sprintf(".%v__%v {\n    \n}\n", class, element)
	_, err = os.Create(filePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, []byte(data), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
