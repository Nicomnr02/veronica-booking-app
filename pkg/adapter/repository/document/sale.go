package document

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	dbConf "pos/internal/config/db"
	dr "pos/pkg/adapter/db"
	"pos/pkg/domain/entity"
	"pos/pkg/shared/util"
)

type Sale struct {
	repo dr.DbDriver
	db   *sql.DB
}

func NewSale(databases dbConf.DatabaseList) *Sale {
	driver, err := dr.NewInstanceDb(databases.HC.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return &Sale{
		repo: driver,
		db:   driver.Db().(*sql.DB),
	}
}

func (a *Sale) CreateSale(req entity.Sale) error {
	sql := `INSERT INTO public.sale ("total_price", "discount", "final_price", "cash", "remaining", "note", "user_id", "items", "created_at", "updated_at") VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	saleID := 0

	err := a.db.QueryRow(sql, req.TotalPrice, req.Discount, req.FinalPrice, req.Cash, req.Remaining, req.Note, req.UserID, req.Items, util.GetCurrentDate(), util.GetCurrentDate()).Scan(&saleID)
	if err != nil {
		panic(err)
	}

	var jsonMapItems []util.PropertyMap
	err = json.Unmarshal([]byte(req.Items), &jsonMapItems)

	for _, items := range jsonMapItems {
		// we are going to make many to many query
		sql := `INSERT INTO public.product_sale ("product_id", "sale_id", "quantity") VALUES($1,$2,$3) RETURNING quantity`

		var i = items["quantity"].(float64)
		var quantity int = int(i)

		var j = items["product_id"].(float64)
		var product_id int = int(j)

		err := a.db.QueryRow(sql, product_id, saleID, quantity).Scan(&quantity)
		fmt.Println(err)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (a *Sale) GetSale(page, limit int) (interface{}, error) {
	offset := (limit*(page-1) + 1) - 1
	sql := `SELECT id, total_price, discount, final_price, cash, remaining, note, user_id, items, created_at, updated_at FROM public.sale order by created_at desc limit $1 offset $2`

	rows, err := a.db.Query(sql, limit, offset)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	sales := []entity.Sale{}

	for rows.Next() {
		sale := entity.Sale{}
		err2 := rows.Scan(&sale.ID, &sale.TotalPrice, &sale.Discount, &sale.FinalPrice, &sale.Cash, &sale.Remaining, &sale.Note, &sale.UserID, &sale.Items, &sale.CreatedAt, &sale.UpdatedAt)
		if err2 != nil {
			return nil, err2
		}
		sales = append(sales, sale)
	}
	return sales, err
}

func (a *Sale) CountSale() (total int, err error) {
	sql := `SELECT count(id) as total
			FROM public.sale`

	err = a.db.QueryRow(sql).Scan(&total)
	return total, err
}
func (a *Sale) GetSaleByID(ID int) (interface{}, error) {
	sql := `SELECT id, total_price, discount, final_price, cash, remaining, note, user_id, items, created_at, updated_at FROM public.sale
			WHERE id = $1`

	sale := entity.Sale{}
	err := a.db.QueryRow(sql, ID).Scan(&sale.ID, &sale.TotalPrice, &sale.Discount, &sale.FinalPrice, &sale.Cash, &sale.Remaining, &sale.Note, &sale.UserID, &sale.Items, &sale.CreatedAt, &sale.UpdatedAt)

	return sale, err
}

func (a *Sale) UpdateSale(req entity.Sale) error {

	sql := `UPDATE public.sale SET "total_price" = $1, "discount"=$2, "final_price"=$3, "cash"=$4, "remaining"=$5, "note"=$6, "user_id"=$7, "items"=$8, "updated_at"=$9 WHERE id = $10`
	stmt, err := a.db.Prepare(sql)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.TotalPrice, req.Discount, req.FinalPrice, req.Cash, req.Remaining, req.Note, req.UserID, req.Items, util.GetCurrentDate(), req.ID)
	if err2 != nil {
		panic(err2)
	}

	var jsonMapItems []util.PropertyMap
	err = json.Unmarshal([]byte(req.Items), &jsonMapItems)

	for _, items := range jsonMapItems {
		// we are going to update many to many query
		sql := `UPDATE public.product_sale SET "quantity"=$1 WHERE sale_id=$2 AND product_id=$3`

		var i = items["quantity"].(float64)
		var quantity int = int(i)

		var j = items["product_id"].(float64)
		var product_id int = int(j)

		stmt, err := a.db.Prepare(sql)

		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err2 := stmt.Exec(quantity, req.ID, product_id)
		if err2 != nil {
			panic(err2)
		}
	}

	return nil
}

func (a *Sale) DeleteSale(req entity.Sale) error {
	sql := `DELETE FROM public.sale WHERE id=$1`
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
