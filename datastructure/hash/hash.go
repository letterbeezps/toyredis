package hash

import (
	"sync"
)

type Hash[T any] struct {
	lock sync.RWMutex
	data map[string]map[string]T
}

func New[T any]() *Hash[T] {
	return &Hash[T]{
		data: map[string]map[string]T{},
	}
}

func (h *Hash[T]) exists(key string) bool {
	_, exist := h.data[key]
	return exist
}

// HKEYS
func (h *Hash[T]) Keys() []string {
	keys := make([]string, 0, len(h.data))

	for k := range h.data {
		keys = append(keys, k)
	}

	return keys
}

// HSET
func (h *Hash[T]) Hset(key, field string, value T) int {
	ret := 0
	if !h.exists(key) {
		h.data[key] = map[string]T{}
	}

	_, exist := h.data[key][field]

	if !exist {
		ret = 1
	}

	h.data[key][field] = value

	return ret
}

// HSETNX
func (h *Hash[T]) Hsetnx(key, field string, value T) int {
	ret := 0
	if !h.exists(key) {
		h.data[key] = map[string]T{}
	}

	if _, exist := h.data[key][field]; !exist {
		ret = 1
		h.data[key][field] = value
	}

	return ret
}

// HGET
func (h *Hash[T]) Hget(key, field string) (T, bool) {
	var value T
	if !h.exists(key) {
		return value, false
	}

	return h.data[key][field], true
}

// HGETALL
func (h *Hash[T]) Hgetall(key string) (keys []string, values []T) {
	if !h.exists(key) {
		return
	}

	for k, v := range h.data[key] {
		keys = append(keys, k)
		values = append(values, v)
	}

	return
}

// HDEL
func (h *Hash[T]) Hdel(key, field string) int {
	if !h.exists(key) {
		return 0
	}

	if _, exist := h.data[key][field]; exist {
		delete(h.data[key], field)
		return 1
	}

	return 0
}

// HEXISTS
func (h *Hash[T]) Hexists(key, field string) (ok bool) {
	if !h.exists(key) {
		return
	}

	if _, exist := h.data[key][field]; exist {
		ok = true
	}

	return
}

// HLEN
func (h *Hash[T]) Hlen(key string) int {
	if !h.exists(key) {
		return 0
	}

	return len(h.data[key])
}

// HKEYS
func (h *Hash[T]) Hkeys(key string) (keys []string) {
	if !h.exists(key) {
		return
	}

	for k := range h.data[key] {
		keys = append(keys, k)
	}

	return
}

// HVALS
func (h *Hash[T]) Hvals(key string) (values []T) {
	if !h.exists(key) {
		return
	}

	for _, v := range h.data[key] {
		values = append(values, v)
	}

	return
}

// HCLEAR
func (h *Hash[T]) Hclear(key string) {
	if !h.exists(key) {
		return
	}

	delete(h.data, key)
}
