package very

import "time"

type ISession interface {
	Expire() time.Duration
	Prefix() string
	New() ISession
}
