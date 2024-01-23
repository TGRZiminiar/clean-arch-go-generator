package server

func (s *server) productService() {
	productRepo := productrepository.NewProductRepository(s.db)
	productUsecase := productusecase.NewProductsUsecase(productRepo)
	productHttpHandler := producthandler.NewProductsHttpHandler(s.cfg, productUsecase)
	
	_ = productHttpHandler
}

