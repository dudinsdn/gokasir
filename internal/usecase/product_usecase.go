package usecase

import (
	"github.com/dudinsdn/gokasir/internal/entity"
	"github.com/dudinsdn/gokasir/internal/repository"
)

type ProductUsecase interface {
	ListProducts(filter string, sort string, page int, pageSize int) ([]entity.Product, int, error)
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: r}
}

func (u *productUsecase) ListProducts(filter string, sort string, page int, pageSize int) ([]entity.Product, int, error) {
	return u.repo.List(filter, sort, page, pageSize)
}
