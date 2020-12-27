package main

import (
	"github.com/prometheus/common/log"
	"app/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd Excute err:%v",err)
	}
}
