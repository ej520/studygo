package main
import (
	"sync"
	"time"
	"fmt"
)

type Bar struct {
	mu       sync.Mutex
	line     int
	prefix   string
	total    int
	width    int
	advance  chan bool
	done     chan bool
	currents map[string]int
	current  int
	before   int
	rate     int
	speed    int
	cost     int
	estimate int
	fast     int
	slow     int
}

var (
	bar1 string
	bar2 string
)

const (
	defaultFast = 20
	defaultSlow = 5
)

func initBar(width int) {
	for i := 0; i < width; i++ {
		bar1 += "="
		bar2 += "-"
	}
}

func NewBar(line int, prefix string, total int) *Bar {
	if total <= 0 {
		return nil
	}

	if line <= 0 {
		gMaxLine++
		line = gMaxLine
	}

	bar := &Bar{
		line:     line,
		prefix:   prefix,
		total:    total,
		fast:     defaultFast,
		slow:     defaultSlow,
		width:    100,
		advance:  make(chan bool),
		done:     make(chan bool),
		currents: make(map[string]int),
	}

	initBar(bar.width)
	go bar.updateCost()
	go bar.run()

	return bar
}

func (b *Bar) SetSpeedSection(fast, slow int) {
	if fast > slow {
		b.fast, b.slow = fast, slow
	} else {
		b.fast, b.slow = slow, fast
	}
}

func (b *Bar) Add() {
	b.current++

	lastRate := b.rate
	lastSpeed := b.speed

	b.count()

	if lastRate != b.rate || lastSpeed != b.speed {
		b.advance <- true
	}

	if b.rate >= 100 {
		close(b.done)
		close(b.advance)
	}
}

func (b *Bar) count() {
	b.mu.Lock()
	now := time.Now()
	nowKey := now.Format("20060102150405")
	befKey := now.Add(time.Minute * -1).Format("20060102150405")
	b.currents[nowKey] = b.current
	if v, ok := b.currents[befKey]; ok {
		b.before = v
	}
	delete(b.currents, befKey)

	b.rate = b.current * 100 / b.total
	if b.cost == 0 {
		b.speed = b.current * 100
	} else if b.before == 0 {
		b.speed = b.current * 100 / b.cost
	} else {
		b.speed = (b.current - b.before) * 100 / 60
	}

	if b.speed != 0 {
		b.estimate = (b.total - b.current) * 100 / b.speed
	}
	b.mu.Unlock()
}

func (b *Bar) updateCost() {
	for {
		select {
		case <-time.After(time.Second):
			b.cost++
			b.count()
			b.advance <- true
		case <-b.done:
			return
		}
	}
}

func (b *Bar) run() {
	for range b.advance {
		printf(b.line, "\r%s", b.barMsg())
	}
}

func (b *Bar) barMsg() string {
	prefix := fmt.Sprintf("%s", b.prefix)
	rate := fmt.Sprintf("%3d%%", b.rate)
	speed := fmt.Sprintf("%3.2fps", 0.01*float64(b.speed))
	cost := b.timeFmt(b.cost)
	estimate := b.timeFmt(b.estimate)
	barLen := b.width - len(prefix) - len(rate) - len(speed) - len(cost) - len(estimate) - 10
	bar1Len := barLen * b.rate / 100
	bar2Len := barLen - bar1Len

	realBar1 := bar1[:bar1Len]
	var realBar2 string
	if bar2Len > 0 {
		realBar2 = ">" + bar2[:bar2Len-1]
	}

	msg := fmt.Sprintf(`%s %s [%s%s] %s %s in: %s`, prefix, rate, realBar1, realBar2, speed, cost, estimate)
	switch {
	case b.speed <= b.slow*100:
		return "\033[0;31m" + msg + "\033[0m"
	case b.speed > b.slow*100 && b.speed < b.fast*100:
		return "\033[0;33m" + msg + "\033[0m"
	default:
		return "\033[0;32m" + msg + "\033[0m"
	}
}

func (b *Bar) timeFmt(cost int) string {
	var h, m, s int
	h = cost / 3600
	m = (cost - h*3600) / 60
	s = cost - h*3600 - m*60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
type Pgbar struct {
	title   string
	maxLine int
	bars    []*Bar
}

func New(title string) *Pgbar {

	p := &Pgbar{
		title:   title,
		maxLine: gMaxLine,
	}

	gMaxLine++
	printf(gMaxLine, "title: %s", title)
	return p
}

func (p *Pgbar) NewBar(prefix string, total int) *Bar {
	gMaxLine++
	return NewBar(gMaxLine, prefix, total)
}



func main() {
	pgb := pgbar.New("多线程进度条")
	pgbar.Println("进度条1")
	b := pgb.NewBar("1st", 20000)
	pgbar.Println("进度条2")
	b2 := pgb.NewBar("2st", 10000)
	pgbar.Println("进度条3")
	b3 := pgb.NewBar("3st", 30000)

	b.SetSpeedSection(900, 100)
	b2.SetSpeedSection(900, 100)
	b3.SetSpeedSection(900, 100)

	pgbar.Println("独立进度条")
	b4 := pgbar.NewBar(0, "4st", 4000)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < 20000; i++ {
			b.Add()
			time.Sleep(time.Second / 2000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			b2.Add()
			time.Sleep(time.Second / 1000)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 30000; i++ {
			b3.Add()
			time.Sleep(time.Second / 3000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 4000; i++ {
			b4.Add()
			time.Sleep(time.Second / 300)
		}
	}()
	wg.Wait()
}