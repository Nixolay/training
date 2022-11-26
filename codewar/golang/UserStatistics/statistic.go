package userstatistics

// Есть последовательность запросов, упорядоченная по времени.
// Запросы бывают двух видов:
//
//    Пользователь user_id сгенерировал событие (нажал на красную кнопку)
//    Посчитать количество пользователей, которые за последние k=5 минут сгенерировали >= limit=1000 событий (нажали на красную кнопку >= 1000 раз).
//
// Необходимо реализовать структуру данных, умеющую эффективно обрабатывать данные запросы.

type UserStatistics struct{}

func NewUserStatistics(k int, limit int) *UserStatistics {
	return nil
}

func (u *UserStatistics) AddEvent(now int, user_id int) {
}

func (u *UserStatistics) GetRobotsCount(now int) int {
	return 0
}
