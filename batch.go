package bbadger

import (
	"github.com/blevesearch/bleve/index/store"
	"github.com/dgraph-io/badger"
)

// Batch implements blevesearch/store/Batch
type Batch struct {
	store    *Store
	merge    *store.EmulatedMerge
	txn      *badger.Txn
	commited bool
}

// Set set the value of the specified key
func (b *Batch) Set(key, val []byte) {
	if err := b.txn.Set(key, val); err == badger.ErrTxnTooBig {
		b.Reset()
		b.txn.Set(key, val)
	}
}

// Delete removes a key and its value
func (b *Batch) Delete(key []byte) {
	b.txn.Delete(key)
}

// Merge applies the default merge policy the specified key's value
func (b *Batch) Merge(key, val []byte) {
	b.merge.Merge(key, val)
}

// Reset resets the current batch
func (b *Batch) Reset() {
	b.txn.Commit()
	b.txn = b.store.db.NewTransaction(true)
	b.merge = store.NewEmulatedMerge(b.store.mo)
}

// Close cleanup the memory
func (b *Batch) Close() error {
	return nil
}
