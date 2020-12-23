package word

import "github.com/spf13/cobra"

var WordCmd = &cobra.Command{
	Use : "word",
	Short:"单词格式转换",
	Long:"支持多种单词格式转换",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
func init(){}
