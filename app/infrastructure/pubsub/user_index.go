package pubsub

import (
	"encoding/json"

	gpubsub "github.com/alash3al/go-pubsub"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/TranTheTuan/authen-go/app/domain/model"
)

type IndexUser struct {
	sub *gpubsub.Subscriber
	db  *gorm.DB
}

func NewIndexUser(db *gorm.DB) *IndexUser {
	return &IndexUser{
		db: db,
	}
}

func (i *IndexUser) Start(broker *gpubsub.Broker, indexName string, topic string) {
	var err error
	i.sub, err = broker.Attach()
	if err != nil {
		log.WithError(err).Error("failed to attach new subscriber")
	}
	broker.Subscribe(i.sub, topic)

	ch1 := i.sub.GetMessages()
	go i.handleEvent(i.sub.GetID(), ch1, indexName)
}

func (i *IndexUser) handleEvent(id string, ch <-chan *gpubsub.Message, indexName string) {
	for {
		if msg, ok := <-ch; ok {
			entry := msg.GetPayload().(*log.Entry)
			if _, ok = entry.Data["request_body"]; ok {
				dataMap := make(map[string]interface{})
				err := json.Unmarshal([]byte(entry.Data["request_body"].(string)), &dataMap)
				if err != nil {
					log.WithError(err).Error("failed to unmarshal request body")
					continue
				}
				indexStore := &model.IndexStore{
					Data: dataMap,
				}
				err = indexStore.MarshalJSON()
				if err != nil {
					log.WithError(err).Error("failed to marshal data json")
					continue
				}
				if err := i.db.Model(model.IndexStore{}).Create(indexStore).Error; err != nil {
					log.WithError(err).Error("failed to insert data to db")
					continue
				}
			}
		}
	}
}
