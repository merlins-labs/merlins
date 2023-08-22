package indexerdb

import "github.com/merlins-labs/merlins-dapp/go/pkg/networks"

type Quest struct {
	ID    string
	Title string
}

type QuestCompletion struct {
	QuestID string          `gorm:"primaryKey"`
	UserID  networks.UserID `gorm:"primaryKey"`

	Quest     *Quest
	Completed bool
}
