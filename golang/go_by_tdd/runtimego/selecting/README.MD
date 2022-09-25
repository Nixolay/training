# SELECT
https://go.dev/src/runtime/select.go 107 строка.

select - это оператор который работает как switch, но для каналов:
```
select {
case msg1 := <- c1: // Получаем сообщение из канала
    fmt.Println("Message 1", msg1)
case msg2, ok := <- c2: // Проверяем был ли получен результат или просто закрыли канал
    if ok {
      fmt.Println("Message 2", msg2)
    }
case <- time.After(time.Second): // По таймауту например можно закрыть соединение с сервером
    fmt.Println("timeout")
default: // Если не один кейс не отработал, можно использовать дефолтный
    fmt.Println("nothing ready")
}
```
В рантайме го представляет функция `selectgo` реализует оператор `select`.

cas0 указывает на массив типа `[ncases]scase`, а order0 указывает на массив типа `[2*ncases]uint16`, где случаи должны быть <= 65536. Оба находятся в стеке goroutine (независимо от любого экранирования в select go).

Для сборок race detector pc0 указывает на массив типа `[ncases]uintptr` (также в стеке); для других сборок он имеет значение nil.

select go возвращает индекс выбранного `scase`, который соответствует порядковому номеру соответствующего вызова `select{recv,send,default}`. Кроме того, если выбранный случай был операцией приема, он сообщает, было ли получено значение.

Select позволяет примерно с одинаковым шансом давать возможность прочитать из них данные.
Проверить это можно с помощью простого теста.
```
func TestGoSelecting(t *testing.T) {
	wg := sync.WaitGroup{}

	for range [1000]int{} {
		wg.Add(1)

		go func() {
			defer wg.Done()

			data := GoSelect()
			for _, item := range data {
				if item < 85 {
					panic(fmt.Sprintf("Item: %d, selection is bad!", item))
				}
			}

			fmt.Println(data)
		}()
	}

	wg.Wait()
}

func GoSelect() []int {
	const lenRange = 1000

	var a, b, c, d, e, f, g int

	ac := make(chan int)
	bc := make(chan int)
	cc := make(chan int)
	dc := make(chan int)
	ec := make(chan int)
	fc := make(chan int)
	gc := make(chan int)

	close(ac)
	close(bc)
	close(cc)
	close(dc)
	close(ec)
	close(fc)
	close(gc)

	for range [lenRange]int{} {
		select {
		case <-ac:
			a++
		case <-bc:
			b++
		case <-cc:
			c++
		case <-dc:
			d++
		case <-ec:
			e++
		case <-fc:
			f++
		case <-gc:
			g++
		}
	}

	return []int{a, b, c, d, e, f, g}
}
```