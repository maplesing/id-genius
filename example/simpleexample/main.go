package main

import (
	"fmt"

	idGen "github.com/maplesing/id-genius"
)

func main() {
	id1, err := generateNewIdByUuid()
	if err != nil {
		fmt.Println("generate new id by UUID failed")
	}
	fmt.Println("id1: ", id1)

	id2, err := generateNewIdBySnowflake()
	if err != nil {
		fmt.Println("generate new id by snowflake failed")
	}
	fmt.Println("id2: ", id2)
}

func generateNewIdByUuid() (string, error) {
	idGen, err := idGen.NewIdGenerator(&idGen.Option{
		IdGeneratorType: idGen.Uuid,
	})
	if err != nil {
		return "", err
	}

	newId, err := idGen.GenerateNewId()
	if err != nil {
		return "", err
	}

	return newId, nil
}

func generateNewIdBySnowflake() (string, error) {
	idGen, err := idGen.NewIdGenerator(&idGen.Option{
		IdGeneratorType: idGen.SnowFlake,
		DatacenterId:    11,
		WorkerId:        30,
	})
	if err != nil {
		return "", err
	}

	newId, err := idGen.GenerateNewId()
	if err != nil {
		return "", err
	}

	return newId, nil
}
