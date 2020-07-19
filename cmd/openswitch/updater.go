package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Updater struct {
	log *logrus.Entry
}

func NewUpdater() *Updater {
	return &Updater{
		log: logrus.WithField("task", "updater"),
	}
}

func (u *Updater) Start(ctx context.Context) {
	u.log.Info("Updater started")
	Fetch(ctx)
	u.log.Info("Updater ended")
}
