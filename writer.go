package bbadger

import (
	"fmt"

	"github.com/blevesearch/bleve/index/store"
	"github.com/dgraph-io/badger"
)

// Writer bleve.search/store/Writer implementation
type Writer struct {
	s *Store
}

// NewBatch creates a new batch
func (w *Writer) NewBatch() store.KVBatch {
	txn := w.s.db.NewTransaction(true)
	return &Batch{
		store: w.s,
		merge: store.NewEmulatedMerge(w.s.mo),
		txn:   txn,
	}
}

// NewBatchEx implements blevesearch.Writer.NewBatchEx
func (w *Writer) NewBatchEx(options store.KVBatchOptions) ([]byte, store.KVBatch, error) {
	return make([]byte, options.TotalBytes), w.NewBatch(), nil
}

// ExecuteBatch implements blevesearch.Writer.ExecuteBatch
func (w *Writer) ExecuteBatch(b store.KVBatch) error {
	batch, ok := b.(*Batch)
	if !ok {
		return fmt.Errorf("wrong type of batch")
	}

	for k, mergeOps := range batch.merge.Merges {
		kb := []byte(k)
		item, err := batch.txn.Get(kb)
		if err != nil && err != badger.ErrKeyNotFound {
			return err
		}
		var v []byte
		if err != badger.ErrKeyNotFound {
			vt, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			v = vt
		}
		mergedVal, fullMergeOk := w.s.mo.FullMerge(kb, v, mergeOps)
		if !fullMergeOk {
			return fmt.Errorf("merge operator returned failure")
		}
		batch.txn.Set(kb, mergedVal)
	}

	return batch.txn.Commit()
}

// Close perform some cleanup operations
func (w *Writer) Close() error {
	return nil
}
