package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}
	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}
	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	return time.Parse("02.01.2006", src)
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	if date.IsZero() {
		return nil, errors.New("неверно указана дата")
	}
	tempMap := make(map[string]time.Duration)
	for _, line := range lines {
		re := regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) (.+)`)
		if re.FindString(line) == "" {
			return nil, errors.New("ошибка в записи данных")
		}
		re = regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) `)
		message := re.Split(line, -1)
		re = regexp.MustCompile(`(\d+:\d+)`)
		timeDur := re.FindAllString(line, -1)
		firstTime, err := time.Parse("15:04", timeDur[0])
		if err != nil {
			return nil, errors.New("ошибка парсинга первого времени")
		}
		secondTime, err := time.Parse("15:04", timeDur[1])
		if err != nil {
			return nil, errors.New("ошибка парсинга второго времени")
		}
		dur := secondTime.Sub(firstTime)
		if dur < 0 {
			return nil, errors.New("неправильно указан диапозон времени")
		}
		if dur == 0 {
			return nil, errors.New("диапозон времени не может быть нулевым")
		}
		_, ok := tempMap[message[1]]
		if !ok {
			tempMap[message[1]] = dur
		} else {
			tempMap[message[1]] += dur
		}
	}
	i := 0
	t := make([]Task, len(tempMap))
	for s, duration := range tempMap {
		t[i].Date = date
		t[i].Dur = duration
		t[i].Title = s
		i++
	}
	return t, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Dur > tasks[j].Dur })
}

// конец решения
// ::footer

func main() {
	page := "25.12.2015\n11:00 - 12:00 first \n12:00 - 12:30 second"

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
