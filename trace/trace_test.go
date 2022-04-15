package trace

import (
	"testing"
	"time"
)

func TestNewQTrace(t *testing.T) {
	tracker := NewQTrace(time.Minute)
	traceAlertEventCh := tracker.Start(100, 1000)
	go func() {
		for v := range traceAlertEventCh {
			_ = v
		}
	}()

	traceOfCh := tracker.TraceOf(&QTraceOfIdListener{
		ListenerCh: make(chan *QTraceOfId),
		Id:         "hello",
	})

	go func() {
		for v := range traceOfCh {
			t.Log(v.Q, v.Qps)
		}
	}()

	for i := 0; i < 10000; i++ {
		go func() {
			for {
				tracker.Trace(&QTraceEntry{
					Id:     "hello",
					On:     time.Now().Unix(),
					Weight: 10,
				})
				time.Sleep(time.Second)
			}
		}()
	}

	c := make(chan struct{})
	<-c
}
