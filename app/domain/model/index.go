package model

import (
	"authen-go/app/domain/common"
	"encoding/json"

	"gorm.io/gorm"
)

type IndexStore struct {
	gorm.Model
	DataJSON common.JSON            `gorm:"type:json"`
	Data     map[string]interface{} `gorm:"-"`
}

func (i *IndexStore) MarshalJSON() error {
	byteVal, err := json.Marshal(&i.Data)
	if err != nil {
		return err
	}
	return i.DataJSON.Scan(byteVal)
}
