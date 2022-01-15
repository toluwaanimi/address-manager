package models

type Address struct {
	BaseModel
	DepositAddress string         `json:"deposit_address"`
	BusinessId     string         `gorm:"type:VARCHAR(30);not null;index:business_id" json:"business_id"`
	Transactions   []Transactions `gorm:"foreignKey:address_id"`
}
