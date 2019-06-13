package gee

import "time"

type ISession interface {
	Save(key string) error
	Get(key string) (ISession, error)
	Expire() time.Duration
	Prefix() string
	New() ISession
}
