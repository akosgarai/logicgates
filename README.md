# Logic gates

Simple application to practice golang with implementing some logic gates

### Useful links about this stuff

- https://hu.wikipedia.org/wiki/Logikai_kapu (hungarian)
- https://en.wikipedia.org/wiki/Logic_gate

## Basic gates

- Not gate

Truth table

|A|Not(A)|
|:-:|:-:|
|1|0|
|0|1|

- And gate

Truth table

|A|B|And(A,B)|
|:-:|:-:|:-:|
|1|1|1|
|1|0|0|
|0|0|0|
|0|1|0|

- Or gate

Truth table

|A|B|Or(A,B)|
|:-:|:-:|:-:|
|1|1|1|
|1|0|1|
|0|0|0|
|0|1|1|

- Nand gate

Truth table

|A|B|Nand(A,B)|
|:-:|:-:|:-:|
|1|1|1|
|1|0|1|
|0|0|0|
|0|1|1|

- Nor gate

Truth table

|A|B|Nor(A,B)|
|:-:|:-:|:-:|
|1|1|0|
|1|0|0|
|0|0|1|
|0|1|0|


## Run tests

```bash
go test -tags="nativ" ./basic/ -v -cover
go test -tags="nand" ./basic/ -v -cover
go test -tags="nor" ./basic/ -v -cover
```

## Run benchmarks

Under `./basic/` directory

```bash
go test -tags="nativ" -bench=.
go test -tags="nand" -bench=.
go test -tags="nor" -bench=.
```
