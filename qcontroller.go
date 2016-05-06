package gormq

type QController interface {
	Routes() map[string]QueueFunc
}
