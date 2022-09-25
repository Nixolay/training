# [Chanel/Каналы](https://go.dev/src/runtime/chan.go "chanel")
[https://go.dev/src/runtime/chan.go](https://go.dev/src/runtime/chan.go "chanel")

Каналы - это структура используемая как основной способ передачи сообщений между каналами.

## Базовая работа с каналами
Каналы бывают буферизированные и не буферизированные.
```
ch := make(chan int, 1) // буферизированный, с размером буфера 1
ch := make(chan int, 0) // не буферизированный, с размером буфера 0
ch := make(chan int)    // не буферизированный, с размером буфера 0
```
Что значит это значит. Буферизированный канал не будет блокировать запись пока его буфер это позволяет.
Не буферизированный канал блокируется при первой же записи.

Пример записи и чтения в каналы:
```
ch := make(chan int, 1)
ch <- 1 // Записываем в канал, что бы тут не упасть на дедлок, необходимо что бы канал был буферизированным.
println(<-ch) // Читаем из канала и отправляем на вывод.
```

Так же для чтения из канала мы можем использовать `range`:
```
ch := make(chan int, 1)
go func(){
  defer close(ch)
  for i := range [10]int{}{
    ch <- i
  }
}()

for e := range ch {
  println(e)
}
```
`range` читает до тех пор пока канал не будет закрыт.

## Направление каналов
По умолчанию каналы двунаправленные, т.е. в них можно писать и читать. Но мы можем задать направление передачи сообщений в канале, сделав его только отправляющим или принимающим.
```
func writer(c chan<- string)
func printer(c <-chan string)
```

# Каналы изнутри
```
type hchan struct {
	qcount   uint           // всего данных в очереди
	dataqsiz uint           // размер циклической очереди
	buf      unsafe.Pointer // указывает на массив элементов
	elemsize uint16
	closed   uint32
	elemtype *_type // тип элемента
	sendx    uint   // индекс элементы массива, который будет отправлен.
	recvx    uint   // индекс массива, куда будет положен новый элемент
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}

type waitq struct {
	first *sudog
	last  *sudog
}
```

рассмотрим какие поля что обозначают в этой структуре:

dataqsize это размер буфера который мы указали при создании канала. elemsize размер одного элемента в канале buf циклическая очередь(циклический буфер) где сохраняются данные. Используется только в буферизированных каналах. closed индикатор закрытого канала. При создании канала это поле 0. После вызова close в это поле устанавливается 1. sendx и recvx это поля для сохранения состояния буфера. Они указывают на позиции в массиве откуда должна происходить отправка или куда должны попадать новые данные. recvq и sendq очереди заблокированных горутин, которые ожидают отправки в канал или чтение из него. lock все отправки и получения должны быть защищены блокировкой. sudog это представление горутины которая стоит в очереди.

## Не буферизированный канал
Давайте создадим канал и отправим в него число
```
ch := make(chan int, 3)
go func(){ch <- 2}()
```
Посмотрим что будет внутри
```
chan int {
    qcount: 0,
    dataqsiz: 0,
    buf: *[0]int [],
    elemsize: 8,
    closed: 0,
    elemtype: *runtime._type {...},
    sendx: 0,
    recvx: 0,
    recvq: waitq<int> { first: nil, lsat: nil },
    sendq: waitq<int> {
        first: *(*sudog<int>)(0xc000074000),
        last: *(*sudog<int>)(0xc000074000),
    },
    lock: runtime.mutex {key:0},
}
```

Как видим буфер у нас длинной `0`, туда мы не можем записать. Это значение сохраняется в структуре sudog. Когда горутина пытается отправить сообщение в канал, но еще нет ни одного получателя. Горутина попадает в список sendq и блокируется.

## Буферизированный канал
Давайте создадим канал 
```
ch := make(chan int, 3)
```
и посмотри что у него будет внутри
```
chan int {
    qcount: 0,
    dataqsiz: 3,
    buf: *[3]int [0,0,0],
    elemsize: 8,
    closed: 0,
    elemtype: *runtime._type {...},
    sendx: 0,
    recvx: 0,
    recvq: waitq<int> { first: nil, lsat: nil },
    sendq: waitq<int> { first: nil, last: nil },
    lock: runtime.mutex {key:0},
}

```

Теперь мы видим что у нас появился буфер из трех элементов и пока этот буфер не будет заполнен, блокировки на запись не произойдет.

# Операции с нулевым каналом
Нулевой канал реагирует на действия с следующим образом:
```
  var c chan int
  v := <-c // блокируется на всегда
  c <- v   // блокируется на всегда
  close(c) // panic
```

Для примера когда читаем из несколько каналов, мы можем организовать завершение цикла чтения:
```
for chanA != nil || chanB != nil {
  select {
  case v, ok := <-chanA:
    if !ok {
      chanA = nil
    }
  case v, ok := <-chanB:
    if !ok {
      chanB = nil
    }
  }
}
```

