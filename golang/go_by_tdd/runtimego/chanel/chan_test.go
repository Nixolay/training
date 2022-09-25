package chanels

import (
	"testing"
)

func TestChanel(t *testing.T) {
	ch := make(chan int, 0)
	ch <- 1       // Записываем в канал, что бы тут не упасть на дедлок, необходимо что бы канал был буферизированным.
	println(<-ch) // Читаем из канала и отправляем на вывод.
}
