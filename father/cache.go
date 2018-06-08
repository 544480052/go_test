package father

type Cacher interface {
	Get(key string) []byte
	Set(key string, val []byte, expire int) error
	Del(key string)
}