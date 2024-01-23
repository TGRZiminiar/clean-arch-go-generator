package producthttphandler

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
