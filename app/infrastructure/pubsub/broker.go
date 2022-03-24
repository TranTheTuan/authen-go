package pubsub

import (
	gpubsub "github.com/alash3al/go-pubsub"
	"gorm.io/gorm"
)

const (
	DefaultEvent = "API_EVENT"
)

var brokerW BrokerWrapper

type BrokerWrapper struct {
	broker *gpubsub.Broker
}

func InitPubSub(db *gorm.DB) {
	brokerW.broker = gpubsub.NewBroker()

	var indexUser = NewIndexUser(db)
	indexUser.Start(brokerW.broker, "user", DefaultEvent)
}

func (b *BrokerWrapper) DispatchEvent(event interface{}) {
	b.broker.Broadcast(event, DefaultEvent)
}
