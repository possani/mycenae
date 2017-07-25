package bcache

import (
	"sync"
	"time"

	"github.com/boltdb/bolt"
	lru "github.com/golang/groupcache/lru"
	"github.com/uol/gobol"
	"github.com/uol/mycenae/lib/tsstats"
)

// BoltCache is a cache using bolt around some DB calls
type BoltCache struct {
	tsCache *lru.Cache
	tsMutex sync.Mutex

	ksCache *lru.Cache
	ksMutex sync.Mutex

	stats *tsstats.StatsTS
	db    *bolt.DB
}

var (
	// BucketKeyspace is the bucket to store keyspace cache data
	BucketKeyspace = []byte("keyspace")

	// BucketTimeseries is the bucket to store timeseries cache data
	BucketTimeseries = []byte("number")

	// BucketText is the bucket to store text timeseries cache data
	BucketText = []byte("text")
)

// NewCache creates a bolt-backed cache
func NewCache(stats *tsstats.StatsTS, path string) (*BoltCache, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return nil, err
	}

	cache := &BoltCache{
		stats: stats,
		db:    db,
	}
	cache.startup()
	go cache.loadData()
	return cache, nil
}

func (bc *BoltCache) startup() error {
	return bc.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(BucketKeyspace); err != nil {
			return err
		}
		if _, err := tx.CreateBucketIfNotExists(BucketTimeseries); err != nil {
			return err
		}
		if _, err := tx.CreateBucketIfNotExists(BucketText); err != nil {
			return err
		}
		return nil
	})
}

func (bc *BoltCache) loadData() {
	bc.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket(BucketTimeseries).ForEach(func(k []byte, v []byte) error {
			bc.tsCache.Add(string(k), nil)
			return nil
		})
	})
}

func (bc *BoltCache) get(bucket, key []byte) ([]byte, gobol.Error) {
	var value []byte
	start := time.Now()
	err := bc.db.View(func(tx *bolt.Tx) error {
		value = tx.Bucket(bucket).Get(key)
		return nil
	})
	if err != nil {
		statsError(bc.stats, "begin", bucket)
		return nil, errPersist("Get", err)
	}
	if value == nil {
		statsNotFound(bc.stats, bucket)
	} else {
		statsSuccess(bc.stats, "get", bucket, time.Since(start))
	}
	return value, nil
}

func (bc *BoltCache) put(bucket, key, value []byte) gobol.Error {
	start := time.Now()
	err := bc.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucket).Put(key, value)
	})
	if err != nil {
		statsError(bc.stats, "put", bucket)
		return errPersist("Put", err)
	}
	statsSuccess(bc.stats, "put", bucket, time.Since(start))
	return nil
}

func (bc *BoltCache) delete(bucket, key []byte) gobol.Error {
	start := time.Now()
	err := bc.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucket).Delete(key)
	})
	if err != nil {
		statsError(bc.stats, "delete", bucket)
		return errPersist("delete", err)
	}
	statsSuccess(bc.stats, "delete", bucket, time.Since(start))
	return nil
}
