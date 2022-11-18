package size

import "pos/pkg/domain/service"

type SizeInputPort interface {
	CreateSize(req service.CreateSizeRequest) error
	GetSize(page, limit int) (interface{}, error)
	GetSizeByID(ID int) (interface{}, error)
	UpdateSize(req service.UpdateSizeRequest) error
	DeleteSize(req service.DeleteSizeRequest) error
}

type SizeOutputPort interface {
	GetResponseSizes(interface{}, int) (interface{}, error)
	GetResponseSize(interface{}) (interface{}, error)
}
