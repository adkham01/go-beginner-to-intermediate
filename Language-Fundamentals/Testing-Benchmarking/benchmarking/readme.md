# Remember to BET
-Benchmark
-Example
-Test

```
-BenchmarkCat(b *testing.B)
-ExampleCat()
-TestCat(t *testing.T)
```

# Commands 
```
go test
go test -v
go test -benchmark .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```