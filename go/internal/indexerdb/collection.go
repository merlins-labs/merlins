package indexerdb

import (
	"time"

	"github.com/merlins-labs/merlins-dapp/go/pkg/networks"
)

type Collection struct {
	// ID is network-dependent
	// Merlins: <id_prefix>-<bech32_mint_contract_address>
	ID networks.CollectionID `gorm:"primaryKey"`

	NetworkId           string
	Name                string
	ImageURI            string
	MaxSupply           int `gorm:"index"`
	SecondaryDuringMint bool
	Paused              bool `gorm:"index"`
	Time                time.Time

	// "has one" relations
	MerlinsCollection *MerlinsCollection

	// "has many" relations
	NFTs []*NFT
}

type MerlinsCollection struct {
	CollectionID        networks.CollectionID `gorm:"index"`
	MintContractAddress string                `gorm:"primaryKey"`
	NFTContractAddress  string
	CreatorAddress      string
	Price               int64
	Denom               string
}
