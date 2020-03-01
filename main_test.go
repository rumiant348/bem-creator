package main

import (
	"fmt"
	"testing"
)

var block = "block"
var blockElement = "block__element"
var blockElementModificator = "block__element_modificator"
var blockElementModificatorValue = "block__element_modificator_value"
var blockModificatorValue = "block_modificator_value"
var blockModificator = "block_modificator"



func Test_getCurrentFolder(t *testing.T) {
	got := getCurrentFolder()
	fmt.Println(got)
}

//func Test_getCurrentFolder(t *testing.T) {
//	got := getCurrentFolder()
//	fmt.Println(got)
//}



