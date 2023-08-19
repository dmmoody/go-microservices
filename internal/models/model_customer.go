package models

type Customer struct {
	CustomerId string `gorm:"primaryKey" json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phoneNumber"`
	Address    string `json:"address"`
}
