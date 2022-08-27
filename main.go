package main

import (
	"fmt"
	"golang-basic/plist/plist-parser"
//	"errors"
)

func main(){
	s := plistparser.GetSystemVersion("SystemVersion.plist")
	fmt.Println(s)
}

