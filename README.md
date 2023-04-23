# IdGenius

IdGenius is a distributed unique ID generator, which can generate distributed unique IDs with different generators.

## The supported generators:

- [x] UUID
- [x] Snowflake

## Examples
### UUID generator
```go
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
```
### Snowflake generator
```go
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
```
## Benchmark tests
```text
BenchmarkUuidGenerator-12                         669319              1643 ns/op              64 B/op          2 allocs/op
BenchmarkSnowFlakeGenerator-12                   4886080               245.4 ns/op            24 B/op          1 allocs/op
BenchmarkUuidGeneratorParallel-12                1816929               785.5 ns/op            64 B/op          2 allocs/op
BenchmarkSnowFlakeGeneratorParallel-12           3339204               356.6 ns/op            24 B/op          1 allocs/op
```