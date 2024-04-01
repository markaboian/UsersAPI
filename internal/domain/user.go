package domain

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"lastName"`
	PlaceOfBirth string `json:"placeOfBirth"`
	DateOfBirth  string `json:"dateOfBirth"`
	Email        string `json:"email"`
}
