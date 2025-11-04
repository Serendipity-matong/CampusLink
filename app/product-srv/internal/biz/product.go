package biz

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Stock       int32
}

type ProductRepo interface {
	CreateProduct(product *Product) error
	GetProduct(id int64) (*Product, error)
	ListProducts(page, pageSize int) ([]*Product, int64, error)
}

type ProductBiz struct {
	repo ProductRepo
}

func NewProductBiz(repo ProductRepo) *ProductBiz {
	return &ProductBiz{repo: repo}
}

