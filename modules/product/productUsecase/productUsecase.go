package productusecase

import productrepository "tgrziminiar/shell-go-clean-arch/modules/product/productRepository"

type (
	ProductUsecaseService interface {
	}

	productUsecase struct {
		productRepo productrepository.ProductRepositoryService
	}
)

func NewProductRepo(productRepo productrepository.ProductRepositoryService) ProductUsecaseService {
	return &productUsecase{
		productRepo: productRepo,
	}
}
