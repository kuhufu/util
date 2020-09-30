package datetime

import (
	"time"
)

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
	return time.Date(year, month, day-1, hour-1, 0, 0, 0, t.Location())
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
	return time.Date(year, month, day-1, hour, min-1, 0, 0, t.Location())
}
