package datetime

import "time"

const (
	DurationDay = time.Hour * 24
)

//流逝的时间
func elapsed(t time.Time) time.Duration {
	now := time.Duration(t.UnixNano()) + zoneOffset(t)
	return now % DurationDay
}

func beginOfToday() time.Time {
	return beginOfCurrDay(time.Now())
}

//一天的开始
func beginOfCurrDay(t time.Time) time.Time {
	return t.Add(-elapsed(t))
}

func beginOfPrevDay(t time.Time) time.Time {
	return t.Add(-elapsed(t) - time.Hour*24)
}

func beginOfNextDay(t time.Time) time.Time {
	return t.Add(-elapsed(t) + time.Hour*24)
}

//时区偏移
func zoneOffset(n time.Time) time.Duration {
	_, offset := n.Zone()
	return time.Duration(offset) * time.Second
}
