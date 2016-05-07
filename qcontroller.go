package redismq

type QController interface {
	Routes() map[string]QueueFunc
}
