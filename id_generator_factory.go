package id_genius

import (
	"fmt"
)

type GeneralIdGeneratorFactory struct {
	option *Option
}

func NewGeneralIdGeneratorFactory(option *Option) *GeneralIdGeneratorFactory {
	return &GeneralIdGeneratorFactory{
		option: option,
	}
}

func (g *GeneralIdGeneratorFactory) CreateIdGenerator() (IdGenerator, error) {
	switch g.option.IdGeneratorType {
	case Uuid:
		return newUuidGenerator(), nil
	case SnowFlake:
		generator, err := getSnowFlakeGenerator(g.option.DatacenterId, g.option.WorkerId)
		if err != nil {
			return nil, err
		}

		return generator, nil
	default:
		return nil, fmt.Errorf("idGeneratorType %v is not found", g.option.IdGeneratorType)
	}
}
