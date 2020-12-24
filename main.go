package main

import (
	"github.com/prometheus/common/log"
	"app/cmd"
)

func main() {
	//p:=flagChild()
	//fmt.Println(p)
	//var name flags.Name
	//	//flag.Var(&name,"name","help")
	//	//flag.Parse()
	//	//fmt.Println(name)

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd Excute err:%v",err)
	}
}
