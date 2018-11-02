Bleve Badger Backend
=====================
> [Blevesearch](https://github.com/blevesearch/bleve) kvstore implementation based on [Badger](https://github.com/dgraph-io/badger) forked from [https://github.com/akhenakh/bleve/tree/badger](https://github.com/akhenakh/bleve/tree/badger) with alot of improvements and fixes.

Usage
==========
> `➜ go get github.com/alash3al/bbadger` .

```go
package main

import (
	"fmt"

	"github.com/alash3al/bbadger"
	"github.com/blevesearch/bleve"
)

func main() {
	// create bleveIndex
	index, err := bbadger.BleveIndex("/tmp/badger/indexName", bleve.NewIndexMapping())

	// or open existing one

	// go on here
	// index.Index(.....)
}

```
