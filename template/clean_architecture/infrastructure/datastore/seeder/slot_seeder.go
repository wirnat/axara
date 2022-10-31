package seeder

import (
	"errors"
	"github.com/google/uuid"
	"gitlab.com/aksaratech/barber-backend/domain/constanta/app"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gitlab.com/aksaratech/barber-backend/infrastructure/clock"
	"gorm.io/gorm"
)

func GenerateSlot(db *gorm.DB) {
	if db.Migrator().HasTable(&model.Slot{}) {
		table := model.Slot{}
		if err := db.Where("interval=?", app.MinutePerSlot).First(&table).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var slots []model.Slot
			times := clock.GenerateTimeInterval(app.MinutePerSlot)
			for i, tt := range times {
				i++
				slot := model.Slot{
					BaseModel: model.BaseModel{
						UUID: uuid.New().String(),
					},
					Time:         tt,
					SerialNumber: i,
					Interval:     app.MinutePerSlot,
				}
				slots = append(slots, slot)
				if err := db.Create(&slot).Error; err != nil {
					panic(err)
				}
			}
		}
	}
}
