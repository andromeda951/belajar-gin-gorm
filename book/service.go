package book

type Service interface {
	Create(bookRequest BookRequest) (Book, error)
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(id int) (Book, error) {
	book, err := s.repository.FindById(id)
	return book, err
}