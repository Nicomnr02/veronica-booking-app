package document

import (
	"database/sql"
	"log"
	dbConf "pos/internal/config/db"
	dr "pos/pkg/adapter/db"
	"pos/pkg/domain/entity"
	"pos/pkg/shared/util"
)

type Category struct {
	repo dr.DbDriver
	db   *sql.DB
}

func NewCategory(databases dbConf.DatabaseList) *Category {
	driver, err := dr.NewInstanceDb(databases.HC.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return &Category{
		repo: driver,
		db:   driver.Db().(*sql.DB),
	}
}

func (a *Category) CreateCategory(req entity.Category) error {
	sql := `INSERT INTO public.category ("category","created_at", "updated_at") VALUES($1,$2,$3) RETURNING id`

	id := 0

	err := a.db.QueryRow(sql, req.Category, util.GetCurrentDate(), util.GetCurrentDate()).Scan(&id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *Category) GetCategory(page, limit int) (interface{}, error) {
	offset := (limit*(page-1) + 1) - 1
	sql := `SELECT * FROM public.category order by created_at desc limit $1 offset $2`

	rows, err := a.db.Query(sql, limit, offset)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	categories := []entity.Category{}

	for rows.Next() {
		category := entity.Category{}
		err2 := rows.Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
		if err2 != nil {
			return nil, err2
		}
		categories = append(categories, category)
	}
	return categories, err
}

func (a *Category) CountCategory() (total int, err error) {
	sql := `SELECT count(id) as total
			FROM public.category`

	err = a.db.QueryRow(sql).Scan(&total)
	return total, err
}
func (a *Category) GetCategoryByID(ID int) (interface{}, error) {
	sql := `SELECT * FROM public.category
			WHERE id = $1`

	category := entity.Category{}
	err := a.db.QueryRow(sql, ID).Scan(&category.ID, &category.Category, &category.CreatedAt, &category.UpdatedAt)

	return category, err
}

func (a *Category) UpdateCategory(req entity.Category) error {
	sql := `UPDATE public.category set "category" = $1, "updated_at" = $2 WHERE id = $3`
	stmt, err := a.db.Prepare(sql)

	// fmt.Println(req.ClassID, req.CategoryID ,req.ID)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Category, util.GetCurrentDate(), req.ID)
	if err2 != nil {
		panic(err2)
	}

	return nil
}

func (a *Category) DeleteCategory(req entity.Category) error {
	sql := `DELETE FROM public.category WHERE id=$1`
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
