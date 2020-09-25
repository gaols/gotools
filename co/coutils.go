package co

import (
	"errors"
	"runtime/debug"
	"sync"
)

// Job abstract a job which can be executed by PoolJobs.
type Job interface {
	// Run run the current job
	Run() (ret interface{}, err error)
}

// PoolJobs execute jobs concurrently with limited concurrency, you can stop on error or ignore errors and proceed on.
func PoolJobs(jobs []Job, concurrency int, stopOnErr bool, retHandler func(job Job, ret interface{}, err error)) {
	if jobs == nil || len(jobs) <= 0 {
		return
	}

	mux := &sync.Mutex{}
	hasErr := false

	sem := make(chan bool, concurrency)
	for _, job := range jobs {
		mux.Lock()
		if hasErr && stopOnErr {
			mux.Unlock()
			break
		}
		mux.Unlock()

		sem <- true
		go func(j Job) {
			defer func() { <-sem }()
			var ret interface{}
			var err error
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = errors.New(string(debug.Stack()))
					}
				}()
				ret, err = j.Run()
			}()

			if err != nil {
				mux.Lock()
				hasErr = true
				mux.Unlock()
			}
			retHandler(j, ret, err)
		}(job)
	}

	for i := 0; i < concurrency; i++ {
		sem <- true
	}
}

// PoolJobsEx get job from channel instead of predefined jobs, jobs channel should be close when no more jobs.
func PoolJobsEx(jobs chan Job, maxConcurrency int, stopOnErr bool, retHandler func(job Job, ret interface{}, err error)) {
	mux := &sync.Mutex{}
	hasErr := false

	sem := make(chan bool, maxConcurrency)
	for job := range jobs {
		mux.Lock()
		if hasErr && stopOnErr {
			mux.Unlock()
			break
		}
		mux.Unlock()

		sem <- true
		go func(j Job) {
			defer func() { <-sem }()
			var ret interface{}
			var err error
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = errors.New(string(debug.Stack()))
					}
				}()
				ret, err = j.Run()
			}()

			if err != nil {
				mux.Lock()
				hasErr = true
				mux.Unlock()
			}
			retHandler(j, ret, err)
		}(job)
	}

	for i := 0; i < maxConcurrency; i++ {
		sem <- true
	}
}
