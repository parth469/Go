package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id         int
	status     bool
	completeBy int
}

func main() {
	workerJob := 4
	totalJobs := 50
	jobQueue := make(chan *Job)
	completeWork := make(chan *Job)

	go func() {
		addTask(jobQueue, totalJobs)
		close(jobQueue)
	}()

	var wg sync.WaitGroup
	wg.Add(workerJob)

	// Start worker pool
	for i := 0; i < workerJob; i++ {
		go workerStart(jobQueue, completeWork, i+1, &wg)
	}

	// Start job result printer
	var done sync.WaitGroup
	done.Add(1)
	go checkJobStats(completeWork, &done)

	wg.Wait()           // wait for all workers to finish
	close(completeWork) // now safe to close result channel
	done.Wait()         // wait for print goroutine
}

// Send jobs to the jobQueue
func addTask(writeJob chan<- *Job, numberOfJob int) {
	for i := 0; i < numberOfJob; i++ {
		jobDetail := &Job{id: i, status: false, completeBy: -1}
		writeJob <- jobDetail
	}
}

// Each worker receives jobs, processes, and sends results
func workerStart(Readjob <-chan *Job, WriteProcessJob chan<- *Job, workerID int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range Readjob {
		job.completeBy = workerID
		fmt.Println("Processing", job.id, "By worker", workerID)
		time.Sleep(300 * time.Millisecond) // Simulated work
		WriteProcessJob <- job
	}
}

// Print processed job results
func checkJobStats(ReadCompletejob <-chan *Job, done *sync.WaitGroup) {
	defer done.Done()
	for job := range ReadCompletejob {
		fmt.Println("✅ Done", job.id, "By worker", job.completeBy)
	}
}
