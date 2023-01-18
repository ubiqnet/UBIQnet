// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package localstore

import (
	"path/filepath"
	"strings"
	"ubiqnet/constants"

	"github.com/syndtr/goleveldb/leveldb/storage"
	"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/syndtr/goleveldb/leveldb"
)

var levelDB *leveldb.DB

func InitDb(path string) error {
	dataPath := filepath.Join(path, "localstore")
	db, err := leveldb.OpenFile(dataPath, nil)
	if err != nil {
		return err
	}
	levelDB = db

	mem, err := leveldb.Open(storage.NewMemStorage(), nil)
	memDB = mem
	return err

}

func Put(key string, value string) {
	levelDB.Put([]byte(key), []byte(value), nil)

}
func Contains(key string) (bool, error) {
	return levelDB.Has([]byte(key), nil)

}

func Get(key string) string {
	data, err := levelDB.Get([]byte(key), nil)
	if err != nil {
		return ""
	}
	result := string(data[:])
	return result
}

func Remove(key string) {
	levelDB.Delete([]byte(key), nil)
}

func FindListRange(prefix string) map[string]string {
	iter := levelDB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	var dict map[string]string = make(map[string]string)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		dict[string(key)] = string(value)
	}
	defer iter.Release()
	return dict
}

func SetNodeType(nodeType string) {
	Put(constants.NodeTypeKey, nodeType)
}

func FindAllHeart() map[string]map[string]string {
	iter := levelDB.NewIterator(util.BytesPrefix([]byte(constants.HeartKeyPrefix)), nil)
	defer iter.Release()
	var total map[string]map[string]string = map[string]map[string]string{}
	var dict map[string]string = make(map[string]string)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		keyArr := strings.Split(string(key), "-")

		dict[string(key)] = string(value)
		if len(dict) >= constants.UploadNum {
			total[keyArr[1]] = dict

		}

	}
	return total
}
