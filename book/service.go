package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(id int, bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository // using Repository interface
}

func NewService(repository Repository) Service { // using Repository interface
	return &service{repository: repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(id int) (Book, error) {
	book, err := s.repository.FindById(id)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(id int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindById(id)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	return newBook, err
}
