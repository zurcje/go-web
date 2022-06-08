package buyer

var b []Buyer
var lastId int64

type Repository interface {
	GetAll() ([]Buyer, error)
	Store(id, cardNumber int64, firstName, lastName string)
	LastId() (int64, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Buyer, error) {
	return b, nil
}

func (r *repository) LastId() (int64, error) {
	return lastId, nil
}

func (r *repository) Store(id, cardNumber int64, firstName, lastName string) (Buyer, error) {
	b := Buyer{id, cardNumber, firstName, lastName}
	b = append(newBuyer, b)
	lastId = b.Id
	return b, nil
}

func NewRepository() Repository {
	return &repository{}
}
