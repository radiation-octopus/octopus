package db

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"sync"
)

var db *leveldb.DB

var once sync.Once

//单例模式
func getInstance() *leveldb.DB {
	once.Do(func() {
		db, _ = leveldb.OpenFile(SaveDbFilePath, nil)
	})
	return db
}

func Start() {
	getInstance()
}

func Stop() {
	getInstance().Close()
}

func getKeyStr(mark string, key string) string {
	return mark + "-" + key
}

//插入
func Insert(mark string, key string, value interface{}) error {
	keyStr := getKeyStr(mark, key)
	valuebyte, error := json.Marshal(&value)
	getInstance().Put([]byte(keyStr), []byte(valuebyte), nil)
	if error != nil {
		return error
	} else {
		return nil
	}
}

//删除
func Delete(mark string, key string) {
	keyStr := getKeyStr(mark, key)
	getInstance().Delete([]byte(keyStr), nil)
}

//查询
func Query(mark string, key string, value interface{}) interface{} {
	keyStr := getKeyStr(mark, key)
	//valueStr := utils.GetInToStr(value)
	valueByte, _ := getInstance().Get([]byte(keyStr), nil)
	json.Unmarshal(valueByte, &value)
	return value
}

//读取mark的所有数据
func QueryMark(mark string) map[string]string {
	iter := db.NewIterator(util.BytesPrefix([]byte(mark+"-")), nil)
	var returnMap map[string]string
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		returnMap[key] = value
	}
	return returnMap
}

//是否包含
func IsHas(mark string, key string) (bool, error) {
	keyStr := getKeyStr(mark, key)
	//valueStr := utils.GetInToStr(value)
	bl, error := getInstance().Has([]byte(keyStr), nil)
	return bl, error
}
