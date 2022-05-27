package models

type Adn struct {
	Secuencia []string `json:"adn"`
	Mutante   bool     `json:"mutante"`
}

type Adns []Adn

// 27017
