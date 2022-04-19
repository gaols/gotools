package trace

import (
	"fmt"
	"github.com/gaols/gotools"
	"math"
	"sync"
	"time"
)

type WinEntry struct {
	Pre  *WinEntry
	Next *WinEntry
	Data TimePoint
}

type TimePoint interface {
	At() int64
}

type SlidingWindow struct {
	Period    time.Duration
	Tolerance time.Duration
	Head      *WinEntry
	Tail      *WinEntry
	Locker    sync.Mutex
}

// t 表示window总时长，tolerance表示最大允许偏差时间，实际上相当总的window时长是t+tolerance
// tolerance的取值一般为添加元素的间隔，例如每20s添加一个元素，那么tolerance则可以取20
func NewSlidingWindow(t time.Duration, tolerance time.Duration) *SlidingWindow {
	win := &SlidingWindow{
		Period:    t,
		Tolerance: tolerance,
	}
	return win
}

func (w *SlidingWindow) Size() int64 {
	w.Locker.Lock()
	defer w.Locker.Unlock()
	p := w.Head
	var c int64
	for p != nil {
		c++
		p = p.Next
	}
	return c
}

func (w *SlidingWindow) print() {
	w.Locker.Lock()
	defer w.Locker.Unlock()
	p := w.Head
	for p != nil {
		fmt.Println(gotools.MustFmtTime(time.Unix(p.Data.At(), 0), "-datetime"))
		p = p.Next
	}
}

func (w *SlidingWindow) HeadData() TimePoint {
	w.Locker.Lock()
	defer w.Locker.Unlock()
	return w.Head.Data
}

// Range return nil means no valid range to calc qps
func (w *SlidingWindow) Range() []TimePoint {
	if w.Head == nil {
		return nil
	}
	pr := int64(w.Period.Seconds())
	t := int64(w.Tolerance.Seconds())
	w.Locker.Lock()
	defer w.Locker.Unlock()
	p := w.Tail
	var minDiff = t
	var minDiffPtr TimePoint
	for p != nil {
		if w.Head.Data.At()-p.Data.At()-pr > t {
			p = p.Pre
		} else {
			diff := int64(math.Abs(float64(w.Head.Data.At() - p.Data.At() - pr)))
			if int64(math.Abs(float64(w.Head.Data.At()-p.Data.At()-pr))) > t {
				break
			} else {
				if diff <= minDiff {
					minDiff = diff
					minDiffPtr = p.Data
				}
			}
			p = p.Pre
		}
	}
	if minDiffPtr != nil {
		return []TimePoint{w.Head.Data, minDiffPtr}
	}

	return nil
}

func (w *SlidingWindow) Add(data TimePoint) {
	w.Locker.Lock()
	defer w.Locker.Unlock()

	node := &WinEntry{
		Pre:  nil,
		Next: nil,
		Data: data,
	}
	if w.Head == nil {
		w.Head = node
		w.Tail = w.Head
		return
	}

	p := w.Head
	for p != nil {
		if p.Data.At() > data.At() {
			p = p.Next
		} else {
			pre := p.Pre
			if pre == nil {
				// add node as new head
				oldHead := w.Head
				w.Head = node
				w.Head.Next = oldHead
				oldHead.Pre = w.Head
			} else {
				node.Pre = pre
				node.Next = p
				p.Pre = node
				pre.Next = node
			}
			return
		}
	}
	// add node as new tail
	oldTail := w.Tail
	oldTail.Next = node
	node.Pre = oldTail
	w.Tail = node
}

func (w *SlidingWindow) Clear() {
	w.Locker.Lock()
	defer w.Locker.Unlock()
	w.Head = nil
	w.Tail = nil
}

// TrimWindow trim the window and return the expired entries
func TrimWindow(w *SlidingWindow) []*WinEntry {
	w.Locker.Lock()
	defer w.Locker.Unlock()
	if w.Head == nil {
		return nil
	}
	ret := make([]*WinEntry, 0, 1)
	p := w.Tail
	oldestValid := time.Now().Unix() - int64((w.Period + w.Tolerance).Seconds()) + 1
	for p != nil {
		if p.Data.At() < oldestValid {
			ret = append(ret, p)
			p = p.Pre
		} else {
			next := p.Next
			w.Tail = p
			p.Next = nil
			if next != nil {
				next.Pre = nil
			}
			return ret
		}
	}
	// clear the window
	w.Tail = nil
	w.Head = nil
	return ret
}
