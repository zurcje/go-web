package buyer

var b []Buyer
var lastId int64

type Repository interface {
	GetAll() ([]Buyer, error)
	Store(id int64, cardNumber int64, firstName string, lastName string) (Buyer, error)
	LastId() (int64, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Buyer, error) {
	return b, nil
}

func (r *repository) LastId() (int64, error) {
	return lastId, nil
}

func (r *repository) Store(id int64, cardNumber int64, firstName string, lastName string) (Buyer, error) {
	newBuyer := Buyer{id, cardNumber, firstName, lastName}
	b = append(b, newBuyer)
	lastId = newBuyer.Id
	return newBuyer, nil
}

func NewRepository() Repository {
	return &repository{}
}
