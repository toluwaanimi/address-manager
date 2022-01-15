package models

type Transactions struct {
	BaseModel
	BusinessId string `gorm:"type:VARCHAR(30);not null;index:business_id" json:"business_id"`
	AddressId  string `gorm:"address_id"`
}
