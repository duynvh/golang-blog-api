package main

import (
	"context"
	"errors"
	"golang-blog-api/component/asyncjob"
	log "golang-blog-api/log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Print("I am job 1")

		return nil
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Print("I am job 2")

		return errors.New("abc")
	})

	// job3 := asyncjob.NewJob(func(ctx context.Context) error {
	// 	time.Sleep(time.Second)
	// 	log.Println("I am job 3")

	// 	return nil
	// })

	// group := asyncjob.NewGroup(true, job1, job2, job3)
	// if err := group.Run(context.Background()); err != nil {
	// 	log.Println(err)
	// }

	if err := job1.Execute(context.Background()); err != nil {
		log.Error("Job 1 err:", err)
		for {
			if err := job1.Retry(context.Background()); err == nil {
				break
			}

			log.Error("Job 1 err", err)
			if job1.State() == asyncjob.StateRetryFailed {
				break
			}
		}
	}

	if err := job2.Execute(context.Background()); err != nil {
		log.Println("Job 2 err:", err)
		for {
			if err := job2.Retry(context.Background()); err == nil {
				break
			}

			log.Println("Job 2 err", err)
			if job2.State() == asyncjob.StateRetryFailed {
				break
			}
		}
	}
}
