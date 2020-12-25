package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	timer "app/internal/time"
)

var calculateTime string
var duration string

var TimeCmd = &cobra.Command{
	Use:"time",
	Short:"时间格式处理",
	Long:"时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var NowTimeCmd = &cobra.Command{
	Use:"now",
	Short:"获取当前时间",
	Long:"获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetTimeNow()
		log.Printf("输出结果: %s,%d",nowTime.Format("2006-01-02 15:04:06"),nowTime.Unix())
	},
}
var CalculateTimeCmd = &cobra.Command{
	Use:"calc",
	Short:"时间格式处理",
	Long:"时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetTimeNow()
		}else{
			var err error
			if !strings.Contains(calculateTime," "){
				layout = "2006-01-02"
			}
			currentTimer,err = time.Parse(layout,calculateTime)
			if err != nil {
				t,_ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t),0)
			}
		}
		calculateTime,err := timer.GetCalculateTime(currentTimer,duration)
		if err != nil {
			log.Fatalf("timer.GetCalcilateTime err:%v",err)
		}
		log.Printf("输出结果:%s,%d",calculateTime.Format(layout),calculateTime)
	},
}

func init(){
	TimeCmd.AddCommand(NowTimeCmd)
	TimeCmd.AddCommand(CalculateTimeCmd)
	CalculateTimeCmd.Flags().StringVarP(&calculateTime,"calculate","c","",`计算时间`)
	CalculateTimeCmd.Flags().StringVarP(&duration,"duration","d","",`持续时间`)
}
