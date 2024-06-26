package object

type CacheObject struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	TTL   int64       `json:"ttl"`
}

func NewCacheObjFromValue(value interface{}) CacheObject {
	return CacheObject{
		Key:   "",
		Value: value,
		TTL:   -1,
	}
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
