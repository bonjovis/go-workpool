# go-workpool
golang workpool 

This work is based on the repo: https://github.com/goinggo/workpool

## Table of contents:
- [Get Started](#get-started)
- [Examples](#examples)


### Get Started
#### Installation

```sh
$ go get github.com/bonjovis/go-workpool
```


#### Examples
```go
import (
  "github.com/bonjovis/go-workpool"
  "log"
  "sync"
  "runtime"
  "strconv"
)

type MyWork struct {
    Name      string "The Name of a person"
    BirthYear int    "The Year the person was born"
}
func NewWorker(name string, year int) *MyWork {
    sw := &MyWork{
        Name:      name,
        BirthYear: year,
    }   
    return sw
}
func (w *MyWork) Work() {
    log.Print(w.Name, " Year:"+strconv.Itoa(w.BirthYear))
}  

func main() {
    start := 0
    end := 20
    size := 1
    w := work.WorkPool(runtime.NumCPU(), logFunc)
    var wg sync.WaitGroup
    wg.Add((end-start)/size + 1)
    for start <= end {
        sw := NewWorker("Name"+strconv.Itoa(start), start)
        go func() {
            w.SetWork(sw)
            wg.Done()
        }()
        start = start + size
    }
    wg.Wait()
}

func logFunc(message string) {
    log.Println(message)
}
```

