package campaign

import "time"

type Campaign struct {
	ID               int    `gorm:"primary_key,type:int"`
	UserID           int    `gorm:"type:int"`
	Name             string `gorm:"size:255"`
	ShortDescription string `gorm:"size:255"`
	Description      string `gorm:"type:text"`
	Perks            string `gorm:"size:255"`
	BackerCount      int    `gorm:"type:int"`
	GoalAmount       int    `gorm:"type:int"`
	CurrentAmount    int    `gorm:"type:int"`
	Slug             string `gorm:"size:255"`
	CampaignImages   []CampaignImage
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CampaignImage struct {
	ID         int    `gorm:"primary_key,type:int"`
	CampaignID int    `gorm:"type:int"`
	FileName   string `gorm:"size:255"`
	IsPrimary  int    `gorm:"type:int"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
