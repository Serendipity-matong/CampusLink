package data

import "github.com/campuslink/campuslink/app/product-srv/internal/biz"

type productRepo struct {
	data *Data
}

func NewProductRepo(data *Data) biz.ProductRepo {
	return &productRepo{
		data: data,
	}
}

func (r *productRepo) CreateProduct(product *biz.Product) error {
	return nil
}

func (r *productRepo) GetProduct(id int64) (*biz.Product, error) {
	return &biz.Product{}, nil
}

func (r *productRepo) ListProducts(page, pageSize int) ([]*biz.Product, int64, error) {
	return nil, 0, nil
}

