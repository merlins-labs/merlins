package indexerdb

import (
	"database/sql"

	"github.com/merlins-labs/merlins-dapp/go/pkg/networks"
)

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type NFT struct {
	// ID is network-dependent
	// Merlins: fury-<bech32_mint_contract_address>-<token_id>
	ID          networks.NFTID
	Name        string
	ImageURI    string
	OwnerID     networks.UserID
	IsListed    bool
	PriceAmount sql.NullString `gorm:"type:numeric"`
	PriceDenom  string
	LockedOn    string

	// "belongs to" relations
	CollectionID networks.CollectionID `gorm:"index"`
	Collection   *Collection

	// "has one" relations
	MerlinsNFT *MerlinsNFT

	// "has many" relations
	Activities []Activity
	Attributes ArrayJSONB `gorm:"type:jsonb;default:'[]'"`
	Burnt      bool
}

type MerlinsNFT struct {
	NFTID   string `gorm:"primaryKey"`
	TokenID string
}
