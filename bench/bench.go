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
	beforeTask func() //任务之前执行
	afterTask  func() //任务之后执行

	concurrency int   //并发度
	total       int   //总任务数
	left        int64 //剩余任务数

	doneC chan struct{}
}

type benchResult struct {
	CountTimes  []time.Duration
	TotalTime   time.Duration
	TotalTask   int
	Concurrency int
	interval    int //间隔单位%
}

func (b *benchResult) Interval(interval int) *benchResult {
	b.interval = interval
	return b
}

func (b *benchResult) String() string {
	step := b.interval
	if step == 0 {
		step = 5
	}

	builder := strings.Builder{}

	n := len(b.CountTimes) - 1

	countTimes := b.countTimesSort()

	for i := step; i <= 100; i += step {
		idx := (n * i) / 100
		builder.WriteString(fmt.Sprintf("%3v%v 小等于 %v\n", i, "%", countTimes[idx]))
	}

	if 100%step != 0 {
		builder.WriteString(fmt.Sprintf("%3v%v 小等于 %v\n", 100, "%", countTimes[len(countTimes)-1]))
	}

	builder.WriteString(fmt.Sprintf("total_time: %v\n", b.TotalTime))
	builder.WriteString(fmt.Sprintf("total_task: %v\n", b.TotalTask))
	builder.WriteString(fmt.Sprintf("concurrency: %v\n", b.Concurrency))

	return builder.String()
}

func (b *benchResult) countTimesSort() []time.Duration {
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

func (b *bench) KeepTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Now().Sub(start)
}

func (b *bench) Do(task func()) *benchResult {

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

	countTimeTask := OneTask
	task = func() {
		elapsed := b.KeepTime(countTimeTask)
		mu.Lock()
		countTimeArr = append(countTimeArr, elapsed)
		mu.Unlock()
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

	return &benchResult{
		CountTimes:  countTimeArr,
		TotalTime:   time.Now().Sub(start),
		TotalTask:   b.total,
		Concurrency: b.concurrency,
	}
}

func Bench() *bench {
	return &bench{
		doneC: make(chan struct{}),
	}
}
