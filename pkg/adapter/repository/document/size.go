package document

import (
	"database/sql"
	"log"
	dbConf "pos/internal/config/db"
	dr "pos/pkg/adapter/db"
	"pos/pkg/domain/entity"
	"pos/pkg/shared/util"
)

type Size struct {
	repo dr.DbDriver
	db   *sql.DB
}

func NewSize(databases dbConf.DatabaseList) *Size {
	driver, err := dr.NewInstanceDb(databases.HC.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect :%v", err)
	}
	return &Size{
		repo: driver,
		db:   driver.Db().(*sql.DB),
	}
}

func (a *Size) CreateSize(req entity.Size) error {
	sql := `INSERT INTO public.size ("size","created_at", "updated_at") VALUES($1,$2,$3) RETURNING id`

	id := 0

	err := a.db.QueryRow(sql, req.Size, util.GetCurrentDate(), util.GetCurrentDate()).Scan(&id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *Size) GetSize(page, limit int) (interface{}, error) {
	offset := (limit*(page-1) + 1) - 1
	sql := `SELECT * FROM public.size order by created_at desc limit $1 offset $2`

	rows, err := a.db.Query(sql, limit, offset)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	sizes := []entity.Size{}

	for rows.Next() {
		size := entity.Size{}
		err2 := rows.Scan(&size.ID, &size.Size, &size.CreatedAt, &size.UpdatedAt)
		if err2 != nil {
			return nil, err2
		}
		sizes = append(sizes, size)
	}
	return sizes, err
}

func (a *Size) CountSize() (total int, err error) {
	sql := `SELECT count(id) as total
			FROM public.size`

	err = a.db.QueryRow(sql).Scan(&total)
	return total, err
}
func (a *Size) GetSizeByID(ID int) (interface{}, error) {
	sql := `SELECT * FROM public.size
			WHERE id = $1`

	size := entity.Size{}
	err := a.db.QueryRow(sql, ID).Scan(&size.ID, &size.Size, &size.CreatedAt, &size.UpdatedAt)

	return size, err
}

func (a *Size) UpdateSize(req entity.Size) error {
	sql := `UPDATE public.size set "size" = $1, "updated_at" = $2 WHERE id = $3`
	stmt, err := a.db.Prepare(sql)

	// fmt.Println(req.ClassID, req.SizeID ,req.ID)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Size, util.GetCurrentDate(), req.ID)
	if err2 != nil {
		panic(err2)
	}

	return nil
}

func (a *Size) DeleteSize(req entity.Size) error {
	sql := `DELETE FROM public.size WHERE id=$1`
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
