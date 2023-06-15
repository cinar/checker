# Optimization Guide

The ```BenchmarkCheck``` function helps profiling the library. Generate a ```profile.out``` my issuing the command below.

```
go test -bench=. -benchmem -cpuprofile profile.out
```

View the analysis through the ```pprof``` by issuing the command below.

```
go tool pprof -http localhost:9000 profile.out
```

Use the web interface for further optimization work.
