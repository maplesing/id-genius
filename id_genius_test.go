package id_genius

import (
	"testing"
)

func TestUuidGenerator(t *testing.T) {
	for i := 0; i < 10; i++ {
		idGen, err := NewIdGenerator(&Option{
			IdGeneratorType: Uuid,
		})
		if err != nil {
			t.Errorf("create id generator failed")
		}

		newId, err := idGen.GenerateNewId()
		if err != nil {
			t.Errorf("generate new id failed")
		}

		t.Logf("generated id %v: %v", i, newId)
	}
}

func TestSnowFlakeGenerator(t *testing.T) {
	for i := 0; i < 10; i++ {
		idGen, err := NewIdGenerator(&Option{
			IdGeneratorType: SnowFlake,
			DatacenterId:    11,
			WorkerId:        30,
		})
		if err != nil {
			t.Errorf("create id generator failed")
		}

		newId, err := idGen.GenerateNewId()
		if err != nil {
			t.Errorf("generate new id failed")
		}

		t.Logf("generated id %v: %v", i, newId)
	}
}

func BenchmarkUuidGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idGen, _ := NewIdGenerator(&Option{
			IdGeneratorType: Uuid,
		})

		idGen.GenerateNewId()
	}
}

func BenchmarkSnowFlakeGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idGen, _ := NewIdGenerator(&Option{
			IdGeneratorType: SnowFlake,
			DatacenterId:    1,
			WorkerId:        10,
		})

		idGen.GenerateNewId()
	}
}

func BenchmarkUuidGeneratorParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idGen, _ := NewIdGenerator(&Option{
				IdGeneratorType: Uuid,
			})

			idGen.GenerateNewId()
		}
	})
}

func BenchmarkSnowFlakeGeneratorParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idGen, _ := NewIdGenerator(&Option{
				IdGeneratorType: SnowFlake,
				DatacenterId:    1,
				WorkerId:        10,
			})

			idGen.GenerateNewId()
		}
	})
}
