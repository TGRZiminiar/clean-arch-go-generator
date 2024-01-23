package productusecase

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
