package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Id          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string     `json:"name"`
	Isbn        string     `json:"isbn"`
	Edition     int        `json:"edition"`
	Publication string     `json:"publication"`
	UserId      uuid.UUID  `json:"userId"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"not null;default:now()"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
