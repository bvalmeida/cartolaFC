package entity

//Assemelha a uma classe na orientação a objetos
type Player struct{
	ID string
	Name string
	Price float64
}

//Assemelha a um construtor
func NewPlayer(id, name string, price float64) *Player{
	return &Player{
		ID: id,
		Name: name,
		Price: price,
	}
}