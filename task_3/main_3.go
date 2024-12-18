package main

type StringIntMap struct {
	mp map[string]int
}

func (mapa *StringIntMap) Add(key string, value int) {
	if mapa.mp == nil {
		mapa.mp = map[string]int{}
	}
	mapa.mp[key] = value
}

func (mapa *StringIntMap) Remove(key string) {
	if mapa.mp == nil {
		return
	}
	delete(mapa.mp, key)
}

func (mapa *StringIntMap) Copy() map[string]int {
	if mapa.mp == nil {
		return map[string]int{}
	}
	newMap := make(map[string]int)
	for k, v := range mapa.mp {
		newMap[k] = v
	}
	return newMap
}

func (mapa *StringIntMap) Exists(key string) bool {
	if mapa.mp == nil {
		return false
	}
	_, ok := mapa.mp[key]
	return ok
}

func (mapa *StringIntMap) Get(key string) (int, bool) {
	if mapa.mp == nil {
		return -1, false
	}
	v, ok := mapa.mp[key]
	return v, ok
}
