package ptime

import "time"

func FormatTimeToDay(timeStamp int64) string {
	str := time.Unix(timeStamp, 0).Format("2006-01-02")
	return str
}

func FormatTimeToSecond(timeStamp int64) string {
	str := time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
	return str
}

func ParseDayToTime(timeStr string) (int64, error) {
	timeStamp, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	return timeStamp.Unix(), nil
}

func ParseSecondToTime(timeStr string) (int64, error) {
	timeStamp, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	return timeStamp.Unix(), nil
}
