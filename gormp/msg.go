package gormq

import "time"

type Msg struct {
	Command string
	ScheduleTo time.Time
	Payload []byte
}
