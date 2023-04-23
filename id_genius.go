package id_genius

func NewIdGenerator(option *Option) (IdGenerator, error) {
	idGeneratorFactory := &GeneralIdGeneratorFactory{
		option: option,
	}

	return idGeneratorFactory.CreateIdGenerator()
}
