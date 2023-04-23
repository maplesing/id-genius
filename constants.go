package id_genius

type IdGeneratorType int32

const (
	Default IdGeneratorType = iota
	Uuid
	SnowFlake
)
