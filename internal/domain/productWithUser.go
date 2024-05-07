package domain

type ProductWithUser struct {
	ProductID        int     `json:"productId"`
	ProductName      string  `json:"productName"`
	ProductPrice     float64 `json:"productPrice"`
	ProductUserId    int     `json:"productUserId"`
	UserId           int     `json:"userId"`
	UserName         string  `json:"userName"`
	UserLastName     string  `json:"userLastName"`
	UserPlaceOfBirth string  `json:"userPlaceOfBirth"`
	UserDateOfBirth  string  `json:"userDateOfBirth"`
	UserEmail        string  `json:"UserEmail"`
}
