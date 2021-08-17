package datetime

import (
	"time"
)

func UnixMilli(milliSeconds int64) time.Time {
	sec := milliSeconds / 1e3
	nsec := milliSeconds % 1e3 * 1e6

	return time.Unix(sec, nsec)
}

//流逝的时间，从 00:00 开始
func Elapsed(t time.Time) time.Duration {
	return t.Sub(BeginOfCurrDay(t))
}

//今天开始的时间
func BeginOfToday() time.Time {
	return BeginOfCurrDay(time.Now())
}

func BeginOfCurrDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func BeginOfNextDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, t.Location())
}

func BeginOfPrevDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day-1, 0, 0, 0, 0, t.Location())
}

func BeginOfCurrHour(t time.Time) time.Time {
	year, month, day := t.Date()
	hour := t.Hour()
	return time.Date(year, month, day, hour, 0, 0, 0, t.Location())
}

func BeginOfNextHour(t time.Time) time.Time {
	year, month, day := t.Date()
	hour := t.Hour()
	return time.Date(year, month, day, hour+1, 0, 0, 0, t.Location())
}

func BeginOfPrevHour(t time.Time) time.Time {
	year, month, day := t.Date()
	hour := t.Hour()
	return time.Date(year, month, day, hour-1, 0, 0, 0, t.Location())
}

func BeginOfCurrMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, min := t.Hour(), t.Minute()
	return time.Date(year, month, day, hour, min, 0, 0, t.Location())
}

func BeginOfNextMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, min := t.Hour(), t.Minute()
	return time.Date(year, month, day, hour, min+1, 0, 0, t.Location())
}

func BeginOfPrevMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, min := t.Hour(), t.Minute()
	return time.Date(year, month, day, hour, min-1, 0, 0, t.Location())
}

/**
获取本周周一的日期
*/
func GetMondayOfCurrentWeek() (weekStartDate time.Time) {
	return GetWeekdayOf(time.Now(), time.Monday)
}

// 这里周日是一周的开始
func GetWeekdayOf(t time.Time, weekday time.Weekday) (weekStartDate time.Time) {
	offset := weekday - t.Weekday()

	weekStartDate = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, int(offset))

	return
}
