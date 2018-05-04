package bbadger

import (
	"github.com/blevesearch/bleve/index/store"
	"github.com/dgraph-io/badger"
)

type Reader struct {
	// you can modify ItrOpts before calling PrefixIterator or PrefixIterator
	// defaulted to badger.DefaultIteratorOptions by store.Reader()
	ItrOpts badger.IteratorOptions
	*badger.Txn
}

func (r *Reader) Get(k []byte) ([]byte, error) {
	item, err := r.Txn.Get(k)
	if err == badger.ErrKeyNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	vs, err := item.Value()
	v := make([]byte, len(vs))
	copy(v, vs)

	return v, err
}

func (r *Reader) MultiGet(keys [][]byte) ([][]byte, error) {
	return store.MultiGet(r, keys)
}

func (r *Reader) PrefixIterator(k []byte) store.KVIterator {
	rv := PrefixIterator{
		iterator: r.Txn.NewIterator(r.ItrOpts),
		prefix:   k[:],
	}
	rv.iterator.Seek(k)
	return &rv
}

func (r *Reader) RangeIterator(start, end []byte) store.KVIterator {
	rv := RangeIterator{
		iterator: r.Txn.NewIterator(r.ItrOpts),
		start:    start[:],
		stop:     end[:],
	}
	rv.iterator.Seek(start)
	return &rv
}

func (r *Reader) Close() error {
	return nil
}
