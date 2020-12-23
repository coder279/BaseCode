package flags

import (
	"errors"
	"flag"
	"fmt"
)

//normal
func Normal() string{
	var name string
	flag.StringVar(&name,"name","test","I will help you")
	flag.Parse()
	fmt.Println(name)
	return name
}
//child flag
func Child()string{
	var name string
	flag.Parse()
	goCmd := flag.NewFlagSet("go",flag.ExitOnError)
	goCmd.StringVar(&name,"name","test","help")
	Args := flag.Args()
	fmt.Println(Args[0])
	_ = goCmd.Parse(Args[1:])
	return name
}

//var
type Name string

func(n *Name)String()string{
	return fmt.Sprint(*n)
}

func (n *Name)Set(value string)error{
	if len(*n) > 0{
		return errors.New("错误的数据")
	}
	*n = Name("bbs" + value)
	return nil
}