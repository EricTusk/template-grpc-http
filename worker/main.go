package worker

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type Worker struct {
	config *Config
}

func NewWorker(cfg *Config) (*Worker, error) {
	if cfg == nil {
		return nil, errors.New("nil config")
	}

	worker := &Worker{
		config: cfg,
	}

	return worker, nil
}

func (w *Worker) Run() error {
	go func() {
		for {
			logrus.Info("worker running....")

			time.Sleep(10 * time.Second)
		}
	}()

	return nil
}

func (w *Worker) Stop() {
	logrus.Info("worker stopping....")

	logrus.Info("worker stopped")
}
