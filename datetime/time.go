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

func BeginOfNextDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d+1, 0, 0, 0, 0, t.Location())
}

func BeginOfPrevDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d-1, 0, 0, 0, 0, t.Location())
}

func BeginOfCurrDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
