package data

type Information struct {
	UserID        string        `json:"user_id" bson:"_id"`
	UserName      UserName      `json:"user_name"`
	Photo         string        `json:"photo"`
	Phone         string        `json:"phone"`
	Birthday      string        `json:"birthday"`
	AddressObject AddressObject `json:"address_object"`
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type UserName struct {
	NickName string `json:"nick_name"`
	Name     Name   `json:"name"`
}
type AddressObject struct {
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
}
