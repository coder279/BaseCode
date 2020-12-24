package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:"",
	Short:"",
	Long:"",
}

func Execute() error{
	return rootCmd.Execute()
}

func init(){
	fmt.Println(1)
	rootCmd.AddCommand(WordCmd)
}

