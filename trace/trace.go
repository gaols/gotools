package trace

import (
	"github.com/gaols/gotools"
	"log"
	"sync"
	"time"
)

type QTrace struct {
	ids            map[string]uint64 // map ip to access count
	TraceDuration  time.Duration
	TraceEntryCh   chan *QTraceEntry
	traceOfCh      chan *QTraceOfIdListener
	traceListeners map[string][]*QTraceOfIdListener
	totalQ         uint64
	// total q of current trace duration window
	winQ           uint64
	win            *SlidingWindow
	lock           sync.Mutex
	expiredCh      chan *QTraceEntry
	PeakWinQps     float64
	PeakWinQpsTime time.Time
}

type QTraceEntry struct {
	Id string
	// trace event occurrence time, you can use time.unix() for this field
	On int64
	// 传入的值应该是你希望的值乘以10，比如一次接口调用，这里应该传1*10
	Weight int
}

func (e *QTraceEntry) At() int64 {
	return e.On
}

func NewQTrace(duration time.Duration) *QTrace {
	return &QTrace{
		ids:            make(map[string]uint64),
		TraceDuration:  duration,
		TraceEntryCh:   make(chan *QTraceEntry, 100),
		traceListeners: make(map[string][]*QTraceOfIdListener),
		traceOfCh:      make(chan *QTraceOfIdListener, 1),
		totalQ:         0,
		win:            NewSlidingWindow(duration, time.Duration(0)),
		expiredCh:      make(chan *QTraceEntry),
		PeakWinQps:     -1,
	}
}

type QTraceOfId struct {
	Q   uint64
	Qps float64
}

type QTraceOfIdListener struct {
	ListenerCh chan *QTraceOfId
	Id         string
}

type QTraceEvent struct {
	Id  string
	Qps float64
}

// Start q trace, winSize is the trace window size
func (t *QTrace) Trace(entry *QTraceEntry) {
	t.TraceEntryCh <- entry
}

func (t *QTrace) TraceOf(listener *QTraceOfIdListener) <-chan *QTraceOfId {
	t.traceOfCh <- listener
	return listener.ListenerCh
}

// Start q trace, winSize is the trace window size
func (t *QTrace) Start(qpsPerIdLimit float64, qpsLimit float64) chan *QTraceEvent {
	eventCh := make(chan *QTraceEvent, 2)
	traceTk := time.NewTicker(time.Minute * 5)
	traceIdTk := time.NewTicker(time.Second)
	go func() {
		defer traceTk.Stop()
		defer traceIdTk.Stop()
		winReady := false
		winSize := int(t.TraceDuration.Seconds())
		startTime := time.Now()
		for {
			if !winReady {
				if time.Now().Sub(startTime).Seconds() >= float64(winSize) {
					winReady = true
				}
			}
			select {
			case entry := <-t.TraceEntryCh:
				t.win.Add(entry)
				t.totalQ += 1
				t.winQ += uint64(entry.Weight)
				v, ok := t.ids[entry.Id]
				if ok {
					v += uint64(entry.Weight)
					t.ids[entry.Id] = v
				} else {
					t.ids[entry.Id] = uint64(entry.Weight)
				}
				if winReady {
					t.calcQpsAndEmitTraceEvent(winSize, qpsLimit, eventCh, entry, qpsPerIdLimit)
				}
			case e := <-t.expiredCh:
				t.winQ -= uint64(e.Weight)
				v := t.ids[e.Id]
				v -= uint64(e.Weight)
				if v == 0 {
					delete(t.ids, e.Id)
				} else {
					t.ids[e.Id] = v
				}
			case l := <-t.traceOfCh:
				listeners := t.traceListeners[l.Id]
				listeners = append(listeners, l)
				t.traceListeners[l.Id] = listeners
			case <-traceIdTk.C:
				if len(t.traceListeners) > 0 {
					for id, idListeners := range t.traceListeners {
						v, ok := t.ids[id]
						if !ok {
							v = 0
						}
						for _, l := range idListeners {
							select {
							case l.ListenerCh <- &QTraceOfId{
								Q:   v / 10,
								Qps: float64(v/10) / float64(winSize),
							}:
							default:
								log.Printf("warning: trace event dropped, id: %s", id)
							}
						}
					}
				}
			case <-traceTk.C:
				peakQpsTime := "-"
				if t.PeakWinQps >= 0 {
					peakQpsTime = gotools.MustFmtTime(t.PeakWinQpsTime, "-datetime")
				}
				log.Printf("[QTrace] winQ/%dm: %d, totalQ: %d, peakWinQps/%dm: %.2f, peakQpsTime: %s", winSize/60, t.winQ/10, t.totalQ, winSize/60, t.PeakWinQps, peakQpsTime)
			}
		}
	}()

	tk := time.NewTicker(time.Second)
	go func() {
		defer tk.Stop()
		for {
			<-tk.C
			expired := TrimWindow(t.win)
			if len(expired) > 0 {
				for _, v := range expired {
					expireEntry := v.Data.(*QTraceEntry)
					t.expiredCh <- expireEntry
				}
			}
		}
	}()

	return eventCh
}

func (t *QTrace) calcQpsAndEmitTraceEvent(winSize int, qpsLimit float64, eventCh chan *QTraceEvent, entry *QTraceEntry, qpsPerIdLimit float64) {
	winQps := float64(t.winQ/10) / float64(winSize)
	if winQps > t.PeakWinQps {
		t.PeakWinQps = winQps
		t.PeakWinQpsTime = time.Now()
	}
	if winQps > qpsLimit {
		eventCh <- &QTraceEvent{
			Id:  "win",
			Qps: winQps,
		}
	}
	idQps := float64(t.ids[entry.Id]/10) / float64(winSize)
	if idQps > qpsPerIdLimit {
		eventCh <- &QTraceEvent{
			Id:  entry.Id,
			Qps: idQps,
		}
	}
}
