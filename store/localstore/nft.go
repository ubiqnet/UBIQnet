package localstore

import (
	"encoding/json"
	"ubiqnet/constants"
)

type NftRecord struct {
	Time   int64  `json:"time"`
	Type   string `json:"type"`
	Source int    `json:"source"`
	Hash   string `json:"hash"`
}

func AddNftRecord(t_time int64, t_type string, source int, hash string) error {
	item := NftRecord{Time: t_time, Type: t_type, Source: source, Hash: hash}
	result := GetNftRecordList()
	result = append(result, item)
	json_str, err := json.Marshal(result)
	if err != nil {
		return err
	}
	Put(constants.NftRecordKey, string(json_str))
	return nil
}

func GetNftRecordList() []NftRecord {
	var list = []NftRecord{}
	result := Get(constants.NftRecordKey)
	json.Unmarshal([]byte(result), &list)
	return list
}
