/*
Package parsefish Напишите простой парсер, который будет анализировать и запускать Deadfish.

У Deadfish есть 4 команды, каждая длиной 1 символ:

     i увеличивает значение (изначально 0)
     d уменьшает значение
     S возводит в квадрат значение
     o выводит значение в возвращаемый массив

     Неверные символы следует игнорировать.

		 Parse ("iiisdoso") == [] int {8, 64}
*/
package parsefish

// Parse by Deadfish.
func Parse(operations string) []int {
	data := 0
	out := make([]int, 0, len(operations))

	for _, operation := range operations {
		switch operation {
		case 'i':
			data++
		case 'd':
			data--
		case 's':
			data *= data
		case 'o':
			out = append(out, data)
		}
	}

	return out
}
