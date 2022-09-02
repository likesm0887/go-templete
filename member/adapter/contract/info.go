package contract

type Information struct {
	UserID        string        `json:"user_id" bson:"_id"`
	UserName      UserName      `json:"UserName"`
	Photo         string        `json:"Photo"`
	Phone         string        `json:"Phone"`
	Birthday      string        `json:"Birthday"`
	AddressObject AddressObject `json:"AddressObject"`
}

type UserName struct {
	NickName string `json:"NickName"`
	Name     Name   `json:"Name"`
}

type Name struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type AddressObject struct {
	Address    string `json:"Address"`
	PostalCode string `json:"PostalCode"`
}
