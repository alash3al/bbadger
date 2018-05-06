package bbadger

import (
	"bytes"

	"gopkg.in/dgraph-io/badger.v1"
)

type RangeIterator struct {
	iterator *badger.Iterator
	start    []byte
	stop     []byte
}

func (i *RangeIterator) Seek(key []byte) {
	if bytes.Compare(key, i.start) < 0 {
		i.iterator.Seek(i.start)
		return
	}
	i.iterator.Seek(key)
}

func (i *RangeIterator) Next() {
	i.iterator.Next()
}

func (i *RangeIterator) Current() ([]byte, []byte, bool) {
	if i.Valid() {
		return i.Key(), i.Value(), true
	}
	return nil, nil, false
}

func (i *RangeIterator) Key() []byte {
	ks := i.iterator.Item().Key()
	k := make([]byte, len(ks))
	copy(k, ks)

	return k
}

func (i *RangeIterator) Value() []byte {
	vs, _ := i.iterator.Item().Value()
	v := make([]byte, len(vs))
	copy(v, vs)

	return v
}

func (i *RangeIterator) Valid() bool {
	if !i.iterator.Valid() {
		return false
	}

	if i.stop == nil || len(i.stop) == 0 {
		return true
	}

	if bytes.Compare(i.stop, i.iterator.Item().Key()) <= 0 {
		return false
	}
	return true
}

func (i *RangeIterator) Close() error {
	i.iterator.Close()
	return nil
}
