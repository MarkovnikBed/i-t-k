package main

import "testing"

func TestAdd(t *testing.T) {
	var mapa StringIntMap
	key := "hours"
	value := 24
	mapa.Add(key, value)
	if mapa.mp[key] != value {
		t.Log("метод не добавил значение")
		t.Fail()
	}

}

func TestRemove(t *testing.T) {
	var mapa StringIntMap
	key := "hours"
	value := 24
	mapa.mp = map[string]int{}
	mapa.mp[key] = value
	mapa.Remove(key)
	if _, ok := mapa.mp[key]; ok {
		t.Log("метод не удалил значение")
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	var mapa StringIntMap
	key1 := "hours"
	value1 := 24
	key2 := "minutes"
	value2 := 3660
	mapa.mp = map[string]int{}
	mapa.mp[key1] = value1
	mapa.mp[key2] = value2
	newMap := mapa.Copy()
	if len(newMap) != len(mapa.mp) {
		t.Log("неправильное копирование, длины мап не совпадают")
		t.Fail()
	} else {
		for k, v := range mapa.mp {
			val, ok := newMap[k]
			if !ok || v != val {
				t.Log("неправильное копирование, не все эелементы есть в новой мапе или значения не соответсвуют ключу")
				t.Fail()
			}
		}
	}

}

func TestExists(t *testing.T) {
	var mapa StringIntMap
	key1 := "hours"
	value1 := 24
	key2 := "minutes"
	value2 := 3660
	mapa.mp = map[string]int{}
	mapa.mp[key1] = value1
	mapa.mp[key2] = value2
	if !mapa.Exists(key1) || !mapa.Exists(key2) {
		t.Log("функция не может вернуть показатель наличия элементов в мапе")
		t.Fail()
	}
	if mapa.Exists("fdsf") {
		t.Log("функция возвращает показатель наличия для несуществуюшего в мапе элемента")
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	var mapa StringIntMap
	key1 := "hours"
	value1 := 24
	key2 := "minutes"
	value2 := 3660
	mapa.mp = map[string]int{}
	mapa.mp[key1] = value1
	mapa.mp[key2] = value2
	v1, ok1 := mapa.Get(key1)
	v2, ok2 := mapa.Get(key2)
	if v1 != value1 || v2 != value2 {
		t.Log("функция не может вернуть значение по ключу в мапе")
		t.Fail()
	}
	if !ok1 || !ok2 {
		t.Log("функция не может вернуть показатель наличия элементов в мапе")
		t.Fail()
	}

	if _, ok := mapa.Get("fsdfs"); ok {
		t.Log("функция возвращает показатель наличия для несуществуюшего в мапе элемента")
		t.Fail()
	}
}
