# MAP

map - в го представляет из себя как хэш-карту.

Хэш-карта - это классическая структура данных, предлагающая O(1) поисковых запросов в среднем и O(n) в худшем случае. То есть, когда все работает хорошо, время выполнения функции map является почти постоянным.

Размер этой константы является частью дизайна hashmap, и точка, в которой карта перемещается от O (1) до O (n) времени доступа, определяется ее хеш-функцией.

Хэш-функция принимает ключ неизвестной длины и возвращает значение с фиксированной длиной.

```hash(key)integer```

Классическая хэш-карта представляет собой массив сегментов, каждый из которых содержит указатель на массив записей ключ/значение. В этом случае наша hashmap имеет восемь сегментов и каждый сегмент может содержать до восьми записей каждый. Использование степеней двойки позволяет использовать дешевые битовые маски и сдвиги, а не дорогостоящее разделение.

По мере добавления записей на карту, предполагая хорошее распределение хэш-функции, сегменты будут заполняться примерно с одинаковой скоростью. Как только количество записей в каждом сегменте превысит некоторый процент от их общего размера, известный как коэффициент загрузки, карта будет расти за счет удвоения количества сегментов и перераспределения записей между ними.

Мы видели, что для реализации hashmap необходимо четыре свойства:
- Для ключа нужна хеш-функция.
- Для сравнения ключей вам нужна функция равенства.
- Вам нужно знать размер ключа.
- Вам нужно знать размер значения, потому что они влияют на размер структуры корзины, которую компилятор должен знать, когда вы переходите или вставляете в эту структуру, как далеко продвигаться в памяти.

map -представляет из себя структуру
https://github.com/golang/go/blob/master/src/runtime/map.go#L116
```
// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
```

map это указатель на hmap в Go (именно он создается при объявлении с помощью var, но не инициализируется, из-за чего падает программа при попытке вставки). Поле buckets — хранилище пар ключ-значение, таких «ведер» несколько, в каждом лежит 8 пар. Сначала в «ведре» лежат слоты для дополнительных битов хэшей. Далее лежат ключи и значения как сначала список всех ключей, потом список всех значений.

По хэш функции определяется в какое «ведро» мы кладем значение, внутри каждого «ведра» может лежать до 8 коллизий, в конце каждого «ведра» есть указатель на дополнительное, если вдруг предыдущее переполнилось.

Основной функционал работы с map(ой):
```
// https://github.com/golang/go/blob/master/src/runtime/map.go#L304
m := make(map[string]int)
// https://github.com/golang/go/blob/master/src/runtime/map.go#L395
v := m["key"]     // runtime.mapaccess1(m, ”key", &v)
// https://github.com/golang/go/blob/master/src/runtime/map.go#L456
v, ok := m["key"] // runtime.mapaccess2(m, ”key”, &v, &ok)
// https://github.com/golang/go/blob/master/src/runtime/map.go#L578
m["key"] = 9001   // runtime.mapassign(m, ”key", 9001)
// https://github.com/golang/go/blob/master/src/runtime/map.go#L695
delete(m, "key")  // runtime.mapdelete(m, “key”)
```

При переборе MAP(ы), мы будем получать разную последовательность:
```
for k, v := range map[int]int{1:1,2:2,3:3,4:4,5:5,6:6} {
	fmt.Printf("key: %d, value: %t\n", k, v)
}
```

Причина кроется в рантайме
```
// https://github.com/golang/go/blob/master/src/runtime/map.go#L847
// mapiterinit initializes the hiter struct used for ranging over maps. 
func mapiterinit(t *maptype, h *hmap, it *hiter) {...
// decide where to start 
r := uintptr(fastrand())
 ... 
it.startBucket = r & bucketMask(h.B)...}
```
Место поиска определяется рандомно.

Мапа не является защищенной для потоков, используйте обертку sync.Mutex или уже готовую реализацию sync.Map