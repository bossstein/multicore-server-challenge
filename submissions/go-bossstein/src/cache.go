package main

type Future[T any] interface {
	Get() T
}

type futureImpl[T any] struct {
	done   chan struct{}
	result T
}

func (f *futureImpl[T]) Get() T {
	<-f.done
	return f.result
}

type req[K comparable, V any] struct {
	key      K
	response chan Future[V]
}

type Cache[K comparable, V any] interface {
	Get(key K) Future[V]
}

type cacheImpl[K comparable, V any] struct {
	requestChannel chan req[K, V]
	getFunction    func(key K) V
}

func (c *cacheImpl[K, V]) Get(key K) Future[V] {
	response := make(chan Future[V], 1)

	c.requestChannel <- req[K, V]{
		key:      key,
		response: response,
	}

	return <-response
}

func (c *cacheImpl[K, V]) Start() {
	store := make(map[K]*futureImpl[V])
	for {
		r := <-c.requestChannel

		future, ok := store[r.key]

		if ok {
			r.response <- future
			continue
		}

		future = &futureImpl[V]{
			done: make(chan struct{}),
		}

		store[r.key] = future
		r.response <- future

		go func() {
			future.result = c.getFunction(r.key)
			close(future.done)
		}()
	}
}

func CreateCache[K comparable, V any](getFunction func(key K) V) Cache[K, V] {
	cache := &cacheImpl[K, V]{
		requestChannel: make(chan req[K, V], 100),
		getFunction:    getFunction,
	}

	go cache.Start()

	return cache
}
