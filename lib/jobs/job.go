package jobs

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers/log"

	"github.com/gocraft/work"
)

var enqueuer *work.Enqueuer

const EnqueuerNamespace = "worker"

type Job interface {
	SetArgs(args map[string]interface{})

	GetArgs() map[string]interface{}

	GetName() string

	Handle() error
}

func Dispatch(job Job) error {
	if enqueuer == nil {
		enqueuer = work.NewEnqueuer(EnqueuerNamespace, database.GetRedisPool())
	}

	workJob, err := enqueuer.Enqueue(job.GetName(), job.GetArgs())
	if err != nil {
		return err
	}

	log.Infof("Enqueued %v job for keyword %v", workJob.Name, workJob.ArgString("keyword"))
	return nil
}
