package models

type Product struct {
	ProductId string `gorm:"primaryKey" json:"productId"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	VendorId  string `json:"vendorId"`
	Vendor    Vendor `gorm:"foreignKey:VendorId" json:"vendor"`
}
