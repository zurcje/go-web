package buyer

import "go-web/internal/domains/file"

var b []Buyer
var lastId int64

type Repository interface {
	GetAll() ([]Buyer, error)
	Store(id int64, cardNumber int64, firstName string, lastName string) (Buyer, error)
	LastId() (int64, error)
}

type repository struct {
	db file.Store
}

func (r *repository) GetAll() ([]Buyer, error) {

	var buyers []Buyer
	if err := r.db.Read(&buyers); err != nil {
		return []Buyer{}, err
	}

	return buyers, nil
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

func NewRepository(db file.Store) Repository {
	return &repository{
		db: db,
	}
}
