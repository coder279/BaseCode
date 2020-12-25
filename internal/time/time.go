package time

import "time"

func GetTimeNow() time.Time{
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time,d string)(time.Time,error){
	duration,err := time.ParseDuration(d)
	if err != nil {
		return time.Time{},err
	}
	return currentTimer.Add(duration),err
}