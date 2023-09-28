package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/dgraph-io/badger"
)

var onceCache sync.Once
var cacheProivder *CacheProvider

func deleteFolder() error {
	err := filepath.Walk("data/cache", func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = os.Remove(ruta)
			if err != nil {
				return err
			}
			fmt.Println("Archivo borrado:", ruta)
		}
		return nil
	})
	return err
}

type CacheProvider struct {
	cache *badger.DB
}

func (r *CacheProvider) InitDB() (*badger.DB, error) {
	deleteFolder()
	opts := badger.DefaultOptions("")
	opts.Dir = "data/cache"
	opts.ValueDir = "data/cache"
	kv, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return kv, nil
}

func (r *CacheProvider) Get(key string) []byte {
	var data []byte
	r.cache.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		item.Value(func(val []byte) error {
			data = val
			return nil
		})
		return err
	})
	return data
}

func (r *CacheProvider) Set(key string, value string) error {
	err := r.cache.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
	return err
}

func GetCacheProvider() *CacheProvider {

	onceCache.Do(func() {
		cacheProivder = new(CacheProvider)
		cache, err := cacheProivder.InitDB()
		fmt.Println(err)
		cacheProivder.cache = cache
	})

	return cacheProivder
}
