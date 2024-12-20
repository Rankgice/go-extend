package ext

type MDict[K comparable, V any] map[K]Vec[V]

func MDict_[K comparable, V any](cap int) MDict[K, V] {
	return make(map[K]Vec[V], cap)
}

// MDictOf 将外部map转为Dict
func MDictOf[K comparable, V any](m map[K]Vec[V]) MDict[K, V] {
	return m
}

func (d MDict[K, V]) ForEach(fn func(KV[K, Vec[V]])) {
	for k, v := range d {
		fn(KV_(k, v))
	}
}

// Len 求出dict长度
func (d MDict[K, V]) Len() int {
	return len(d)
}

// Empty 判断dict是否为空
func (d MDict[K, V]) Empty() bool {
	return d.Len() == 0
}

// Load 判断key是否存在,并求出对应的值
func (d MDict[K, V]) Load(key K) Opt[Vec[V]] {
	v, b := d[key]
	return Opt_(v, b)
}

// Store 添加键值对
func (d MDict[K, V]) Store(key K, value V) {
	d[key] = append(d[key], value)
}

// MStore 添加键值对
func (d MDict[K, V]) MStore(key K, values ...V) {
	d[key] = append(d[key], values...)
}

// LoadOrMStore 向Dict中添加键值对，如果key存在，则直接返回
func (d MDict[K, V]) LoadOrMStore(key K, values ...V) Vec[V] {
	v, b := d[key]
	if !b {
		d[key] = values
	}
	return v
}

// LoadAndDelete 通过key删除键值对，并且返回v和b，如果key不存在则返回nil，false
func (d MDict[K, V]) LoadAndDelete(key K) Vec[V] {
	v, b := d[key]
	if b {
		delete(d, key)
	}
	return v
}

// Delete 删除键值对
func (d MDict[K, V]) Delete(key K) {
	delete(d, key)
}

// ToVec 将dict转为Vec
func (d MDict[K, V]) ToVec() Vec[KV[K, Vec[V]]] {
	vec := Vec_[KV[K, Vec[V]]](d.Len())
	for k, v := range d {
		vec.Append(KV_(k, v))
	}
	return vec
}

// Keys 获取所有的key
func (d MDict[K, V]) Keys() Vec[K] {
	vec := Vec_[K](d.Len())
	for k := range d {
		vec.Append(k)
	}
	return vec
}

// Values 获取所有的Values
func (d MDict[K, V]) Values() Vec[Vec[V]] {
	vec := Vec_[Vec[V]](d.Len())
	for _, v := range d {
		vec.Append(v)
	}
	return vec
}

func (d MDict[K, V]) Clear() {
	clear(d)
}

func (d MDict[K, V]) AppendSelf(kv KV[K, Vec[V]]) MDict[K, V] {
	d[kv.K] = kv.V
	return d
}
