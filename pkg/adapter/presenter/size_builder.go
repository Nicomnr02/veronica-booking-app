package presenter

import (
	"pos/pkg/domain/entity"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type SizeBuilder struct{}

func (*SizeBuilder) GetResponseSizes(s interface{}, totalData int) (interface{}, error) {
	var sizes []entity.Sizes
	err := mapstructure.Decode(s, &sizes)
	if err != nil {
		panic(err)
	}

	return service.SizeResponses{
		Sizes: sizes,
		Total: totalData,
	}, nil
}

func (*SizeBuilder) GetResponseSize(u interface{}) (interface{}, error) {
	var size entity.Size
	err := mapstructure.Decode(u, &size)
	if err != nil {
		return nil, err
	}

	return service.SizeResponse{
		Size: size,
	}, nil
}
