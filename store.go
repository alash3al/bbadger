// MIT LICENSE
//
//  Copyright (c) 2017 Fabrice Aneche
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package bbadger

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve/index/store"
	"github.com/blevesearch/bleve/registry"
	"gopkg.in/dgraph-io/badger.v1"
)

const (
	Name = "badger"
)

type Store struct {
	path string
	db   *badger.DB
	mo   store.MergeOperator
}

func New(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	path, ok := config["path"].(string)
	if !ok {
		return nil, fmt.Errorf("must specify path")
	}
	if path == "" {
		return nil, os.ErrInvalid
	}

	opt := badger.DefaultOptions
	opt.Dir = path
	opt.ValueDir = path

	if cdir, ok := config["create_if_missing"].(bool); ok && cdir {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, os.FileMode(0700))
			if err != nil {
				return nil, err
			}
		}
	}

	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	rv := Store{
		path: path,
		db:   db,
		mo:   mo,
	}
	return &rv, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Reader() (store.KVReader, error) {
	return &Reader{
		ItrOpts: badger.DefaultIteratorOptions,
		Txn:     s.db.NewTransaction(false),
	}, nil
}

func (s *Store) Writer() (store.KVWriter, error) {
	return &Writer{
		s: s,
	}, nil
}

func init() {
	registry.RegisterKVStore(Name, New)
}
