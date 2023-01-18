// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package localstore

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var memDB *leveldb.DB

func MemPut(key string, value string) {
	_ = memDB.Put([]byte(key), []byte(value), nil)

}
func MemContains(key string) bool {
	b, _ := memDB.Has([]byte(key), nil)
	return b

}

func MemGet(key string) string {
	data, err := memDB.Get([]byte(key), nil)
	if err != nil {
		return ""
	}
	result := string(data[:])
	return result
}

func MemRemove(key string) {
	_ = memDB.Delete([]byte(key), nil)
}

func MemFindListRange(prefix string) map[string]string {
	iter := memDB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	var dict map[string]string = make(map[string]string)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		dict[string(key)] = string(value)
	}
	defer iter.Release()
	return dict
}
