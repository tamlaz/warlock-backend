package cron

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/go-co-op/gocron/v2"
)

func CleanUpQaJob() {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		log.Printf("Failed to start cron scheduler")
		panic(err)
	}

	job, err := scheduler.NewJob(
		gocron.DurationJob(
			5*time.Minute,
		),
		gocron.NewTask(
			RemoveQaOlderThanFiveMinutes,
		),
		gocron.WithName("job: removing QA records older than 5 minutes"),
	)

	if err != nil {
		panic(err)
	}

	log.Println(job.ID())

	scheduler.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Interrupt signal received. Exiting...")
		_ = scheduler.Shutdown()
		os.Exit(0)
	}()

	for {

	}
}

func RemoveQaOlderThanFiveMinutes() {
	cutoff := time.Now().Add(-5 * time.Minute)
	if err := config.DB.Where("created_at < ?", cutoff).Delete(&models.Qa{}).Error; err != nil {
		panic(err)
	}
}
