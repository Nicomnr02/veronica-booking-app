package document

import (
	"database/sql"
	"log"
	dbConf "pos/internal/config/db"
	dr "pos/pkg/adapter/db"
	"pos/pkg/domain/entity"
	"pos/pkg/shared/util"
)

type Product struct {
	repo dr.DbDriver
	db   *sql.DB
}

// UpdateProductStock implements repository.ProductRepository
func (*Product) UpdateProductStock(req entity.Product) error {
	panic("unimplemented")
}

func NewProduct(databases dbConf.DatabaseList) *Product {
	driver, err := dr.NewInstanceDb(databases.HC.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return &Product{
		repo: driver,
		db:   driver.Db().(*sql.DB),
	}
}

func (a *Product) CreateProduct(req entity.Product) error {
	sql := `INSERT INTO public.product ("category_id","size_id","barcode","name","price","stock","detail_product","created_at", "updated_at") VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

	id := 0

	err := a.db.QueryRow(sql, req.CategoryID, req.SizeID, req.Barcode, req.Name, req.Price, 0, req.DetailProduct, util.GetCurrentDate(), util.GetCurrentDate()).Scan(&id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *Product) GetProduct(page, limit int) (interface{}, error) {
	offset := (limit*(page-1) + 1) - 1
	sql := `SELECT * FROM public.product order by created_at desc limit $1 offset $2`

	rows, err := a.db.Query(sql, limit, offset)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	products := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}
		err2 := rows.Scan(&product.ID, &product.CategoryID, &product.SizeID, &product.Barcode, &product.Name, &product.Price, &product.Stock, &product.DetailProduct, &product.CreatedAt, &product.UpdatedAt)
		if err2 != nil {
			return nil, err2
		}
		products = append(products, product)
	}
	return products, err
}

func (a *Product) CountProduct() (total int, err error) {
	sql := `SELECT count(id) as total
			FROM public.product`

	err = a.db.QueryRow(sql).Scan(&total)
	return total, err
}

func (a *Product) GetProductByID(ID int) (interface{}, error) {
	sql := `SELECT * FROM public.product
			WHERE id = $1`

	product := entity.Product{}
	err := a.db.QueryRow(sql, ID).Scan(&product.ID, &product.CategoryID, &product.SizeID, &product.Barcode, &product.Name, &product.Price, &product.Stock, &product.DetailProduct, &product.CreatedAt, &product.UpdatedAt)

	return product, err
}

func (a *Product) UpdateProduct(req entity.Product) error {
	sql := `UPDATE public.product set "category_id" = $1, "size_id"=$2, "barcode"=$3, "name"=$4, "price"=$5,  "detail_product"=$6, "updated_at" = $7, "stock" = $8 WHERE id = $9`
	stmt, err := a.db.Prepare(sql)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.CategoryID, req.SizeID, req.Barcode, req.Name, req.Price, req.DetailProduct, util.GetCurrentDate(), req.ID)
	if err2 != nil {
		panic(err2)
	}

	return nil
}

func (a *Product) DeleteProduct(req entity.Product) error {
	sql := `DELETE FROM public.product WHERE id=$1`
	stmt, err := a.db.Prepare(sql)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.ID)
	if err2 != nil {
		panic(err2)
	}
	return nil

}
