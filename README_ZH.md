# IdGenius
IdGenius 是一个分布式全局ID生成器，可以使用不同的方法生成，比如：UUID，雪花算法。

## 支持的生成方法

- [x] UUID
- [x] 雪花算法

## 使用方式
### UUID
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
### 雪花算法
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
## Benchmark 测试
```text
BenchmarkUuidGenerator-12                         669319              1643 ns/op              64 B/op          2 allocs/op
BenchmarkSnowFlakeGenerator-12                   4886080               245.4 ns/op            24 B/op          1 allocs/op
BenchmarkUuidGeneratorParallel-12                1816929               785.5 ns/op            64 B/op          2 allocs/op
BenchmarkSnowFlakeGeneratorParallel-12           3339204               356.6 ns/op            24 B/op          1 allocs/op
```
## 参考
* [twitter snowflake](https://github.com/twitter-archive/snowflake)