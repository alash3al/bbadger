package bbadger

import (
	"bytes"

	"gopkg.in/dgraph-io/badger.v1"
)

type PrefixIterator struct {
	iterator *badger.Iterator
	prefix   []byte
}

func (i *PrefixIterator) Seek(key []byte) {
	if bytes.Compare(key, i.prefix) < 0 {
		i.iterator.Seek(i.prefix)
		return
	}
	i.iterator.Seek(key)
}

func (i *PrefixIterator) Next() {
	i.iterator.Next()
}

func (i *PrefixIterator) Current() ([]byte, []byte, bool) {
	if i.Valid() {
		return i.Key(), i.Value(), true
	}
	return nil, nil, false
}

func (i *PrefixIterator) Key() []byte {
	return i.iterator.Item().KeyCopy(nil)
}

func (i *PrefixIterator) Value() []byte {
	v, _ := i.iterator.Item().ValueCopy(nil)

	return v
}

func (i *PrefixIterator) Valid() bool {
	return i.iterator.ValidForPrefix(i.prefix)
}

func (i *PrefixIterator) Close() error {
	i.iterator.Close()
	return nil
}
