package bbadger

import (
	"github.com/blevesearch/bleve/index/store"
	"gopkg.in/dgraph-io/badger.v1"
)

type Batch struct {
	store *Store
	merge *store.EmulatedMerge
	*badger.Txn
}

func (b *Batch) Set(key, val []byte) {
	keyc := make([]byte, len(key))
	copy(keyc, key)

	valc := make([]byte, len(val))
	copy(valc, val)

	b.Txn.Set(keyc, valc)
}

func (b *Batch) Delete(key []byte) {
	keyc := make([]byte, len(key))
	copy(keyc, key)

	b.Txn.Delete(keyc)
}

func (b *Batch) Merge(key, val []byte) {
	b.merge.Merge(key, val)
}

func (b *Batch) Reset() {
	b.merge = store.NewEmulatedMerge(b.store.mo)
	b.Txn = b.store.db.NewTransaction(true)
}

func (b *Batch) Close() error {
	b.merge = nil
	b.Txn.Discard()
	return nil
}
