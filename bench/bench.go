package bench

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type bench struct {
	countTime  bool   //记录任务时间
	beforeTask func() //任务之前执行
	afterTask  func() //任务之后执行

	concurrency int   //并发度
	total       int   //总任务数
	left        int64 //剩余任务数

	doneC chan struct{}
}

type benchResult struct {
	owner      *bench
	CountTimes []time.Duration
	Total      time.Duration
}

func (b *benchResult) ToString(step int) string {
	builder := strings.Builder{}

	n := len(b.CountTimes) - 1

	countTimes := b.CountTimesSort()

	for i := step; i <= 100; i += step {
		idx := (n * i) / 100
		builder.WriteString(fmt.Sprintf("%3v%v 小等于 %v\n", i, "%", countTimes[idx]))
	}

	if 100%step != 0 {
		builder.WriteString(fmt.Sprintf("%3v%v 小等于 %v\n", 100, "%", countTimes[len(countTimes)-1]))
	}

	builder.WriteString(fmt.Sprintf("total_time: %v\n", b.Total))
	builder.WriteString(fmt.Sprintf("total_task: %v\n", b.owner.total))
	builder.WriteString(fmt.Sprintf("concurrency: %v\n", b.owner.concurrency))

	return builder.String()
}

func (b *benchResult) CountTimesSort() []time.Duration {
	cpy := append(b.CountTimes[:0:0], b.CountTimes...)

	sort.Slice(cpy, func(i, j int) bool {
		if cpy[i] < cpy[j] {
			return true
		}
		return false
	})

	return cpy
}

func (b *bench) Concurrency(num int) *bench {
	b.concurrency = num
	return b
}

func (b *bench) Total(num int) *bench {
	b.total = num
	b.left = int64(num)
	return b
}

func (b *bench) Before(f func()) *bench {
	b.beforeTask = f
	return b
}

func (b *bench) After(f func()) *bench {
	b.afterTask = f
	return b
}

func (b *bench) CountTime() *bench {
	b.countTime = true
	return b
}

func (b *bench) KeepTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Now().Sub(start)
}

func (b *bench) Do(task func()) benchResult {

	start := time.Now()

	mu := sync.Mutex{}

	OneTask := task
	task = func() {
		if b.beforeTask != nil {
			b.beforeTask()
		}
		OneTask()
		if b.afterTask != nil {
			b.afterTask()
		}
	}

	var countTimeArr []time.Duration

	if b.countTime {
		countTimeTask := OneTask
		task = func() {
			elapsed := b.KeepTime(countTimeTask)
			mu.Lock()
			countTimeArr = append(countTimeArr, elapsed)
			mu.Unlock()
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(b.concurrency)
	for i := 0; i < b.concurrency; i++ {
		go func() {
			defer wg.Done()
			for {
				v := atomic.AddInt64(&b.left, -1)
				if v < 0 {
					return
				}
				task()
			}
		}()
	}

	wg.Wait()

	return benchResult{
		owner:      b,
		CountTimes: countTimeArr,
		Total:      time.Now().Sub(start),
	}
}

func Bench() *bench {
	return &bench{
		doneC: make(chan struct{}),
	}
}
