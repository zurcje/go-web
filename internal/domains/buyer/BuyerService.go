package buyer

type Service interface {
	GetAll() ([]Buyer, error)
	Store(cardNumber int64, firstName string, lastName string) (Buyer, error)
	//LastId() int64
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Buyer, error) {
	b, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *service) Store(cardNumber int64, firstName string, lastName string) (Buyer, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return Buyer{}, err
	}
	lastId++

	buyer, err := s.repository.Store(lastId, cardNumber, firstName, lastName)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
