package timetools

import "time"

func GetNowDateStr(layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	t := time.Now().Format(layout)
	return t
}

// 时间戳转年月日 时分秒
func DateFormat(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02 15:04")
}

//时间戳转年月日
func DateFormatYmd(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02")
}

//获取当前年月
func DateYmFormat() string {
	tm := time.Now()
	return tm.Format("2006-01")
}

//获取年月日时分秒（字符串类型）
func DateNowFormatStr() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}

//时间戳
func DateUnix() int {
	t := time.Now().Local().Unix()
	return int(t)
}

//获取年月日时分秒(time类型)
func DateNowFormat() time.Time {
	tm := time.Now()
	return tm
}

//获取第几周
func DateWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}

//获取年、月、日
func DateYMD() (int, int, int) {
	year, month, day := DateYmdInts()
	return year, month, day
}

// 获取年月日
func DateYmdFormat() string {
	tm := time.Now()
	return tm.Format("2006-01-02")
}

// 获取日期的年月日
func DateYmdInts() (int, int, int) {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	return year, int(month), day
}
