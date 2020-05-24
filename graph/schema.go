package graph

type Management struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Zip     string `json:"zip"`
	Address string `json:"address"`
}

type Manager struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// Managemnet    Management `json:"managemnet" gorm:"foreignkey:ManagemnetsID"`
	ManagemnetsID uint
}
