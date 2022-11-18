package document

import (
	"database/sql"
	"log"
	dbConf "pos/internal/config/db"
	dr "pos/pkg/adapter/db"
	"pos/pkg/domain/entity"
	"pos/pkg/shared/util"
)

type Stock struct {
	repo dr.DbDriver
	db   *sql.DB
}

func NewStock(databases dbConf.DatabaseList) *Stock {
	driver, err := dr.NewInstanceDb(databases.HC.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return &Stock{
		repo: driver,
		db:   driver.Db().(*sql.DB),
	}
}

func (a *Stock) CreateStock(req entity.Stock) error {
	sql := `INSERT INTO public.stock ("product_id","detail","quantity","created_at", "updated_at") VALUES($1,$2,$3,$4,$5) RETURNING id`

	id := 0

	err := a.db.QueryRow(sql, req.ProductID, req.Detail, req.Quantity, util.GetCurrentDate(), util.GetCurrentDate()).Scan(&id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *Stock) GetStock(page, limit int) (interface{}, error) {
	offset := (limit*(page-1) + 1) - 1
	sql := `SELECT * FROM public.stock order by created_at desc limit $1 offset $2`

	rows, err := a.db.Query(sql, limit, offset)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	stocks := []entity.Stock{}

	for rows.Next() {
		stock := entity.Stock{}
		err2 := rows.Scan(&stock.ID, &stock.ProductID, &stock.Detail, &stock.Quantity, &stock.CreatedAt, &stock.UpdatedAt)
		if err2 != nil {
			return nil, err2
		}
		stocks = append(stocks, stock)
	}
	return stocks, err
}

func (a *Stock) CountStock() (total int, err error) {
	sql := `SELECT count(id) as total
			FROM public.stock`

	err = a.db.QueryRow(sql).Scan(&total)
	return total, err
}

func (a *Stock) GetStockByID(ID int) (interface{}, error) {
	sql := `SELECT * FROM public.stock
			WHERE id = $1`

	stock := entity.Stock{}
	err := a.db.QueryRow(sql, ID).Scan(&stock.ID, &stock.ProductID, &stock.Detail, &stock.Quantity, &stock.CreatedAt, &stock.UpdatedAt)

	return stock, err
}

func (a *Stock) UpdateStock(req entity.Stock) error {
	sql := `UPDATE public.stock set "product_id" = $1, "detail"=$2, "quantity"=$3, "updated_at" = $4 WHERE id = $5`
	stmt, err := a.db.Prepare(sql)

	// fmt.Println(req.ClassID, req.StockID ,req.ID)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.ProductID, req.Detail, req.Quantity, util.GetCurrentDate(), req.ID)
	if err2 != nil {
		panic(err2)
	}

	return nil
}

func (a *Stock) DeleteStock(req entity.Stock) error {
	sql := `DELETE FROM public.stock WHERE id=$1`
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
