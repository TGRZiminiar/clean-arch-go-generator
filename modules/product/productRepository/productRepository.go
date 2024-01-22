package productrepository

type (
	ProductRepositoryService interface {

	}

	productRepository struct {
		db     any
	}
)

func NewProductRepository(db any) ProductRepositoryService {
	return &productRepository{
		db:     db,
	}
}
