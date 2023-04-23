package id_genius

type IdGeneratorFactory interface {
	CreateIdGenerator() (IdGenerator, error)
}

type IdGenerator interface {
	GenerateNewId() (string, error)
}
