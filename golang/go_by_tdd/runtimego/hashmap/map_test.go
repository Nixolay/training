package hashmap

/* func TestHashMap(t *testing.T) {
	hm := New(10)

	hm.Set("hello", "world")
	hm.Set("world", "hello")
} */

/* func New(size int) HashMap {
	return HashMap{
		data: make([]ListElement, size),
	}
}

type HashMap struct {
	data []ListElement
}

func (hm *HashMap) Set(key, value string) {
	hm.data[getStringHash(key)].Set(key, value)
}

func (hm HashMap) Get(key string) (string, bool) {
	value := hm.data[getStringHash(key)].Del(key)
	return value, value != ""
}

func (hm *HashMap) Del(key string) {
	hm.data[getStringHash(key)].Del(key)
}

type ListElement struct {
	kvList []KeyValue
}

type KeyValue struct{ key, value string }

func (le *ListElement) Set(key, value string) {
	var kv *KeyValue
	for i := range le.kvList {
		if le.kvList[i].key == key {
			kv = &le.kvList[i]
		}
	}

	if kv != nil {
		kv.value = value
	} else {
		le.kvList = append(le.kvList, KeyValue{key: key, value: value})
	}
}

func (le *ListElement) Get(key string) string {
	for i := range le.kvList {
		if le.kvList[i].key == key {
			return le.kvList[i].value
		}
	}

	return ""
}

func (le *ListElement) Del(key string) {
	i := 0
	for i = range le.kvList {
		if le.kvList[i].key == key {
			break
		}
	}

	le.kvList = append(le.kvList[:i], le.kvList[i:]...)
} */
