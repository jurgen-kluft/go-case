# Console Application:

## Command-Line arguments:
    
https://github.com/codegangsta/cli
https://github.com/alecthomas/kingpin
https://github.com/codegangsta/cli
    
## Console:
https://github.com/alecthomas/colour (colouring of the output)


# Hashing:

## Skein:
    
https://github.com/yann2192/GoKeeper/blob/master/src/crypto/skein/skein.go
https://github.com/whyrusleeping/GoSkein

# Configuration file formats:

https://github.com/spf13/viper     (YAML, JSON, TOML, HCL)
https://github.com/vaughan0/go-ini (INI)
https://github.com/go-ini/ini      (INI)


# Compression

## LZ4:

https://github.com/pierrec/lz4
https://github.com/bkaradzic/go-lz4

## Snappy:

https://github.com/golang/snappy


# Logging

## log15

https://www.progville.com/go/log15-powerful-logging-golang/

## Piping command-line tools (interesting)

https://blog.gopheracademy.com/advent-2015/composable-command-line-tools/


```

// Work-Pool (http://play.golang.org/p/KUVQEEEdfV)

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type PoolProcessable interface {
	Process() (out PoolProcessable)
}

type group []PoolProcessable

type Pool struct {
	Work chan PoolProcessable
	todo chan group
	Res  chan PoolProcessable
}

func newPool(workers, groupFactor int) *Pool {
	var wg sync.WaitGroup
	p := &Pool{
		Work: make(chan PoolProcessable, 100),
		todo: make(chan group, 100),
		Res:  make(chan PoolProcessable, 100),
	}

	// launch workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for g := range p.todo {
				for _, thing := range g {
					p.Res <- thing.Process()
				}
			}
			wg.Done()
		}()
	}

	// launch grouper
	wg.Add(1)
	go func() {
		var g group
		for w := range p.Work {
			g = append(g, w)
			if len(g) == groupFactor {
				p.todo <- g
				g = nil
			}
		}
		p.todo <- g
		close(p.todo)
		wg.Done()
	}()

	// launch finisher
	go func() { wg.Wait(); close(p.Res) }()

	return p
}

type ppString string

func (pp ppString) Process() (out PoolProcessable) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	return ppString(string(pp) + " processed")
}

func main() {
	pool := newPool(4, 10)

	// Send in the work; do it a a goroutine so that
	// if we block on send, the result reader can still run
	// and drain the queues.
	go func() {
		for i := 0; i < 1000; i++ {
			pool.Work <- ppString(fmt.Sprintf("line %d", i))
		}
		close(pool.Work)
	}()

	// Read all the results.
	for res := range pool.Res {
	        pps := res.(ppString)
		fmt.Println(string(pps))
	}
}
```


