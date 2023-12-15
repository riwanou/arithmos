# Advanced Algorithms Project

In data, add `cles_alea` and `Shakespeare`. \
The project with everything is available on github here: https://github.com/riwanou/arithmos.

## Run Bench and Plots

```bash
go test ./lib -bench . | tee >(awk "/Benchmark/") > bench_output
python3 visu_bench.py
```

## Run tests

```bash
go test ./lib -v
```

