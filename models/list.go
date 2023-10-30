package models

import (
	"time"
)

type List struct {
	SubList
	SubLists []SubList `gorm:"foreignKey:ParentID" json:"subLists,omitempty"`
}

type SubList struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title" validate:"required,max=100"`
	Description string    `gorm:"type:text;not null" json:"description" validate:"required,max=1000"`
	Priority    int       `gorm:"not null" json:"priority"`
	ParentID    *uint     `json:"parent_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`

	Attachments []*Attachment `gorm:"foreignKey:ListID" json:"attachments"`
}

type Attachment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	ListID    uint      `gorm:"not null" json:"list_id"`
	Filename  string    `gorm:"not null" json:"filename"`
	Filepath  string    `gorm:"not null" json:"filepath"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`
}

func (List) TableName() string {
	return "list"
}

func (SubList) TableName() string {
	return "list"
}

func (Attachment) TableName() string {
	return "attachment"
}
