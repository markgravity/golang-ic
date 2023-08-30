package main

import (
	"github.com/markgravity/golang-ic/bootstrap"
	"github.com/markgravity/golang-ic/lib/jobs"
	"os"
	"os/signal"
	"syscall"

	"github.com/markgravity/golang-ic/database"

	"github.com/gocraft/work"
)

type Context struct{}

func init() {
	bootstrap.LoadEnv()
	bootstrap.InitDatabase(database.GetDatabaseURL())
	database.SetupRedisDB()
}

func main() {
	pool := work.NewWorkerPool(Context{}, 5, jobs.EnqueuerNamespace, database.GetRedisPool())

	crawl := jobs.Crawl{}
	pool.JobWithOptions(crawl.GetName(), work.JobOptions{MaxFails: 3}, func(job *work.Job) error {
		crawl.SetArgs(job.Args)
		return crawl.Handle()
	})

	pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	pool.Stop()
}
