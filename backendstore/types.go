package backendstore


type Store interface {
	SetIfNotExists(key, token string, ttl int, ch chan bool)
	DelIfKeyHasVal(key, token string, ch chan struct{})
}