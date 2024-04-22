package object

type CacheObject struct {
	Key   string
	Value interface{}
	TTL   int64
}

func NewCacheObjFromParams(key string, value interface{}, ttl int64) CacheObject {
	return CacheObject{
		Key:   key,
		Value: value,
		TTL:   ttl,
	}
}

func NewEmptyCacheObj() CacheObject {
	return CacheObject{
		Key:   "",
		Value: nil,
		TTL:   -1,
	}
}
