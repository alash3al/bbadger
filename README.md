Bleve Badger Backend
=====================
> [Blevesearch](https://github.com/blevesearch/bleve) kvstore implementation based on [Badger](https://github.com/dgraph-io/badger) forked from [https://github.com/akhenakh/bleve/tree/badger](https://github.com/akhenakh/bleve/tree/badger) with some fixed issues.

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
	index, err := bleve.NewUsing("indexName", bleve.NewIndexMapping(), bleve.Config.DefaultIndexType, bbadger.Name, map[string]interface{}{
		"path": "/tmp/badger",
	})

	// go on here
	// index.Index(.....)
}

```
