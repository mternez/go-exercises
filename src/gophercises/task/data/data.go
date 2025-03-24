package data

import "github.com/boltdb/bolt"

type Value interface {
	GetKey() string
}

type Repository[V Value] struct {
	Name        string
	Deserialize func(data []byte) (V, error)
	Serialize   func(value V) ([]byte, error)
}

func (r *Repository[V]) getBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	b, err := tx.CreateBucketIfNotExists([]byte(r.Name))
	if err != nil {
		return tx.Bucket([]byte(r.Name)), nil
	}
	return b, err
}

func (r *Repository[V]) Get(db *bolt.DB, key string) (V, error) {
	var value V
	err := db.View(func(tx *bolt.Tx) error {
		bucket, err := r.getBucket(tx)
		if err != nil {
			return err
		}
		data := bucket.Get([]byte(key))
		value, err = r.Deserialize(data)
		return err
	})
	return value, err
}
func (r *Repository[V]) Put(db *bolt.DB, value V) error {
	valueBytes, err := r.Serialize(value)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, er := r.getBucket(tx)
		if er != nil {
			return er
		}
		return bucket.Put([]byte(value.GetKey()), valueBytes)
	})
	return err
}
func (r *Repository[V]) Values(db *bolt.DB) ([]V, error) {
	values := make([]V, 0)
	err := db.View(func(tx *bolt.Tx) error {
		bucket, err := r.getBucket(tx)
		if err != nil {
			return err
		}
		return bucket.ForEach(func(k, v []byte) error {
			deserializedValue, er := r.Deserialize(v)
			if er != nil {
				return er
			}
			values = append(values, deserializedValue)
			return err
		})
	})
	return values, err
}
