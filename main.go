package main

import (
	"app/flags"
	"flag"
	"fmt"
)
func flagNomal()string{
	return flags.Normal()
}

func flagChild()string{
	p:= flags.Child()
	return p
}

func main() {
	//p:=flagChild()
	//fmt.Println(p)
	var name flags.Name
	flag.Var(&name,"name","help")
	flag.Parse()
	fmt.Println(name)
}
