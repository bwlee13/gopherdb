package cache

type Cache interface {
	Put()
	Get()
	Add()
	Delete()
	Flush()
	CountKeys()
}
