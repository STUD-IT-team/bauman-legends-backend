package cache

import "sync"

// ICache
//
// Интерфейс кэша. Содержит методы:
//   - Put (создание/изменение записи по ключу)
//   - Delete (удаление записи по ключу)
//   - Find (поиск записи по ключу)
type ICache[K comparable, V any] interface {
	Put(id K, value V)
	Delete(id K)
	Find(id K) *V
}

// Cache
//
// Общая структура кэша:
//   - mutex - для осуществления конкурентной записи
//   - value - карта с ключом K и значением V
type Cache[K comparable, V any] struct {
	mutex sync.RWMutex
	value map[K]V
}

// NewCache
//
// Инициализация объекта кэша
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		value: make(map[K]V),
	}
}

// Put
//
// Создание новой (или изменение существующей) записи по ключу
func (c *Cache[K, V]) Put(id K, value V) {
	defer c.mutex.Unlock()

	c.mutex.Lock()
	c.value[id] = value
}

// Delete
//
// Удаление записи по ключу
func (c *Cache[K, V]) Delete(id K) {
	defer c.mutex.Unlock()

	c.mutex.Lock()
	delete(c.value, id)
}

// Find
//
// Поиск записи по ключу.
// Возвращает указатель на запись
// (при отсутствии записи в карте возвращается nil)
func (c *Cache[K, V]) Find(id K) *V {
	defer c.mutex.RUnlock()

	c.mutex.RLock()
	value, exists := c.value[id]
	if exists {
		return &value
	}
	return nil
}
