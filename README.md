# Advanced Algorithms Project

## Run Bench and Plots

```bash
go test ./lib -bench . | tee >(awk "/Benchmark/") > bench_output
python3 visu_bench.py
```

## Run tests

```bash
go test ./lib -v
```

