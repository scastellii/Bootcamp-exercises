package products

import (
	"go-bases/02-go-web/01-API-tt/api-food/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByPrice(price float64) ([]domain.Product, error)
	GetById(id int) (*domain.Product, error)
	Store(product *domain.Product) (*domain.Product, error)
	Update(id int, product *domain.Product) (*domain.Product, error)
	Delete(id int) (*domain.Product, error)
}

func NewServiceController(repo *RepositoryController) *ServiceController {
	return &ServiceController{
		repository: repo,
	}
}

type ServiceController struct {
	repository Repository
}

func (s *ServiceController) GetAll() ([]domain.Product, error) {
	all, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return all, err
}

func (s *ServiceController) GetByPrice(price float64) ([]domain.Product, error) {
	product, err := s.repository.GetByPrice(price)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (s *ServiceController) GetById(id int) (*domain.Product, error) {
	product, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (s *ServiceController) Store(product *domain.Product) (*domain.Product, error) {
	prd, err := s.repository.Store(product)
	if err != nil {
		return nil, err
	}
	return &prd, nil
}

func (s *ServiceController) Update(id int, product *domain.Product) (*domain.Product, error) {
	updateProduct, err := s.repository.Update(id, product)
	if err != nil {
		return nil, err
	}
	return &updateProduct, nil
}

func (s *ServiceController) Delete(id int) (*domain.Product, error) {
	product, err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
