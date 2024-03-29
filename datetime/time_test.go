package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestBeginOfToday(t *testing.T) {
	fmt.Println(BeginOfToday())
}

func TestBeginOfCurrDay(t *testing.T) {
	fmt.Println(BeginOfCurrDay(time.Now()))
}

func TestBeginOfPrevDay(t *testing.T) {
	fmt.Println(BeginOfPrevDay(time.Now()))
}

func TestBeginOfNextDay(t *testing.T) {
	fmt.Println(BeginOfNextDay(time.Now()))
}

func BenchmarkBeginOfCurrDay(b *testing.B) {
	now := time.Now()
	for i := 0; i < b.N; i++ {
		BeginOfCurrDay(now)
	}
}

func TestGetMondayOfCurrentWeek(t *testing.T) {
	fmt.Println(GetMondayOfCurrentWeek())
}

func TestGetWeekdayOf(t *testing.T) {
	sunday := time.Now().Add(time.Hour * 24 * 3)
	fmt.Println(GetWeekdayOf(sunday, time.Sunday))
}
