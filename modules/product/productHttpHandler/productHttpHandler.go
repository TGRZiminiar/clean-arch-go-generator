package producthttphandler

import productusecase "tgrziminiar/shell-go-clean-arch/modules/product/productUsecase"

type (
	ProductHttpHandlerService interface {
	}

	productHttpHandler struct {
		productUsecase productusecase.ProductUsecaseService
	}
)

func NewProductHttpHandler(productUsecase productusecase.ProductUsecaseService) ProductHttpHandlerService {
	return &productHttpHandler{
		productUsecase: productUsecase,
	}
}
