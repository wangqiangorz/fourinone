package model

type Pikachu struct {
	Ketchup
	status int
}

func NewPikachu() *Pikachu {
	pikachu := new(Pikachu)
	pikachu.status = NOTREADY
	pikachu.Init()
	return pikachu
}

func (pikachu *Pikachu) SetStatus(status int) {
	pikachu.status = status
}

func (pikachu *Pikachu) GetStatus() int {
	return pikachu.status
}
