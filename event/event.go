package event

import "github.com/lorezi/golang-cqrs-pattern/schema"

type EventStore interface {
	Close()
	PublishMeowCreated(meow schema.Meow) error
	SubscribeMeowCreated() (<-chan MeowCreatedMessage, error)
	OnMeowCreated(f func(MeowCreatedMessage)) error
}
