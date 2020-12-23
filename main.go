package main

import (
	"app/flags"
	"flag"
	"fmt"
	"github.com/spf13/cobra"
)


func main() {
	//p:=flagChild()
	//fmt.Println(p)
	var name flags.Name
	flag.Var(&name,"name","help")
	flag.Parse()
	fmt.Println(name)
}
