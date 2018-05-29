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
	if err := b.Txn.Set(key, val); err == badger.ErrTxnTooBig {
		b.Txn.Commit(nil)
		b.Txn = b.store.db.NewTransaction(true)
		b.Txn.Set(key, val)
	}
}

func (b *Batch) Delete(key []byte) {
	b.Txn.Delete(key)
}

func (b *Batch) Merge(key, val []byte) {
	b.merge.Merge(key, val)
}

func (b *Batch) Reset() {
	b.Txn.Commit(nil)
	b.Txn = b.store.db.NewTransaction(true)
	b.merge = store.NewEmulatedMerge(b.store.mo)
}

func (b *Batch) Close() error {
	b.merge = nil
	b.Txn.Commit(nil)
	return nil
}
