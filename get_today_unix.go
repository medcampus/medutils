package medutils

import "time"

func GetTodayUnix() int64 {
	year, month, date := time.Now().Date()
	local, _ := time.LoadLocation("")
	return time.Date(year, month, date, 0, 0, 0, 0, local).UTC().Unix()
}
