package repository

import "pos/pkg/domain/entity"

type SizeRepository interface {
	CreateSize(req entity.Size) error
	GetSize(page, limit int) (interface{}, error)
	GetSizeByID(ID int) (interface{}, error)
	UpdateSize(req entity.Size) error
	DeleteSize(req entity.Size) error
	CountSize() (total int, err error)
}
