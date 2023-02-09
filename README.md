# Go 1.20 new features

Quick review of Go 1.20 features

- [Go 1.20 new features](#go-120-new-features)
  - [Profile Guided Optimization](#profile-guided-optimization)
    - [Merging different profiles](#merging-different-profiles)
  - [Integration test coverage](#integration-test-coverage)
  - [Array to slice conversion](#array-to-slice-conversion)
  - [New errors wrapping](#new-errors-wrapping)

## Profile Guided Optimization

Profile guided optimization is used to generate more optimized code from using CPU profiles from captured running system.

1. Add profiling to your program `defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()`
2. Build your program as usual
3. Run your program with under load
4. Build your program with flag `-pgo <path/to/profile>`
5. Enjoy

### Merging different profiles

You can run profiling of your code multiple times. Then you should merge your profiles `go tool pprof -proto a.pprof b.pprof > merged.pprof`

## Integration test coverage

Now you can check coverage of integration tests of go programs. To try this just do following steps:

1. Build program with coverage flag `go build -cover ...`
2. Run program with set env `GOCOVERDIR`
3. Optionally merge results `go tool covdata merge -i=<directories> -o merged`

## Array to slice conversion

You can now convert arrays to 

```go
s := make([]byte, 2, 4)
a0 := [0]byte(s)
a1 := [1]byte(s[1:]) // a1[0] == s[1]
```

## New errors wrapping

Errors now can wrap multiple errors

```go
err1 := fmt.Errorf("first sample error")
err2 := fmt.Errorf("second sample error")
err3 := fmt.Errorf("third sample error")

complexErr1 := fmt.Errorf("first complex error %w, %w", err1, err2)
complexErr2 := fmt.Errorf("first complex error %w, %w", complexErr1, err3)
```
