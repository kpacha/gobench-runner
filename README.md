# gobench-runner
run golang benchmarks and then complete and store the reports

## Installation

```
go get github.com/kpacha/gobench-runner
go install github.com/kpacha/gobench-runner
```

## Usage

The binary flags are properly documented

```
$ gobench-runner -h
Usage of gobench-runner:
  -b string
    	branch under test (default "master")
  -c string
    	sha of the commit
  -d string
    	date of the commit
  -nr
    	no-recursive
  -o string
    	path of the folder where the results are stored (default "/tmp/")
```

A quick usage example

```
$ gobench-runner -b flatmap_formatter -c 01c6e63 -d "2019-03-23T17:36:0+0100" github.com/devopsfaith/krakend/sd
2019/03/23 22:07:59 benchmarking package github.com/devopsfaith/krakend/sd
2019/03/23 22:08:21 benchmark results stored at /tmp/bench_result_01c6e63_github.com___devopsfaith___krakend___sd
# Binaries for programs and plugins
2019/03/23 22:08:21 benchmarking package github.com/devopsfaith/krakend/sd/dnssrv
2019/03/23 22:08:21 no benchmarks found in package github.com/devopsfaith/krakend/sd/dnssrv
2019/03/23 22:08:21 done
```

The reports are ready to be sent to the [golang perf service](https://perf.golang.org/) using the [`benchsave` cmd](https://godoc.org/golang.org/x/perf/cmd/benchsave)

```
$ ./benchsave -server http://localhost:8080 -v /tmp/bench_result_01c6e63_github.com___devopsfaith___krakend___sd
```