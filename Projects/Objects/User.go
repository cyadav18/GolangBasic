package Objects

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	PersonID  int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Address   Address `gorm:"foreignKey:UserName"`
}

type Address struct {
	gorm.Model
	AddressID int `gorm:"primaryKey"`
	Number    string
	Area      string
	City      string
	State     string
}

func (customer *Customer) TableName() string {
	return "customer"
}

func (address *Address) TableName() string {
	return "customer_address"
}

