package models

type User struct {
	ID       uint64   `json:"id" bson:"id"`
	Name     string   `json:"name" bson:"user_name"`
	LastName string   `json:"last_name" bson:"user_last_name"`
	Age      uint64   `json:"age" bson:"user_age"`
	Address  *Address `json:"address" bson:"user_address"`
}

type Address struct {
	ID       uint64 `json:"id" bson:"id"`
	City     string `json:"city" bson:"city"`
	Street   string `json:"street" bson:"street"`
	Postcode string `json:"postcode" bson:"postcode"`
}
