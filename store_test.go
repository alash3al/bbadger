package bbadger

import (
	"os"
	"testing"

	"log"

	"github.com/blevesearch/bleve/index/store"
	"github.com/blevesearch/bleve/index/store/test"
)

func open(t *testing.T, mo store.MergeOperator) store.KVStore {
	rv, err := New(mo, map[string]interface{}{
		"path":              "testbadger",
		"create_if_missing": true,
	})
	if err != nil {
		t.Fatal(err)
	}
	return rv
}

func cleanup(t *testing.T, s store.KVStore) {
	err := s.Close()
	if err != nil {
		log.Printf("%+v\n", err)
		t.Fatal(err)
	}
	err = os.RemoveAll("testbadger")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBadgerKVCrud(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestKVCrud(t, s)
}

func TestBadgerReaderIsolation(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestReaderIsolation(t, s)
}

func TestBadgerReaderOwnsGetBytes(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestReaderOwnsGetBytes(t, s)
}

func TestBadgerWriterOwnsBytes(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestWriterOwnsBytes(t, s)
}

func TestBadgerPrefixIterator(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestPrefixIterator(t, s)
}

func TestBadgerPrefixIteratorSeek(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestPrefixIteratorSeek(t, s)
}

func TestBadgerRangeIterator(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestRangeIterator(t, s)
}

func TestBadgerRangeIteratorSeek(t *testing.T) {
	s := open(t, nil)
	defer cleanup(t, s)
	test.CommonTestRangeIteratorSeek(t, s)
}

func TestBadgerMerge(t *testing.T) {
	s := open(t, &test.TestMergeCounter{})
	defer cleanup(t, s)
	test.CommonTestMerge(t, s)
}
