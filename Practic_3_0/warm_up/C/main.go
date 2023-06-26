package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// начало решения

// SQLMap представляет карту, которая хранится в SQL-базе данных
type SQLMap struct{}

// NewSQLMap создает новую SQL-карту в указанной базе
func NewSQLMap(db *sql.DB) (*SQLMap, error) {
	return &SQLMap{}, nil
}

// Get возвращает значение для указанного ключа.
// Если такого ключа нет - возвращает ошибку sql.ErrNoRows.
func (m *SQLMap) Get(key string) (any, error) {
	return nil, nil
}

// Set устанавливает значение для указанного ключа.
// Если такой ключ уже есть - затирает старое значение (это не считается ошибкой).
func (m *SQLMap) Set(key string, val any) error {
	return nil
}

// Delete удаляет запись карты с указанным ключом.
// Если такого ключа нет - ничего не делает (это не считается ошибкой).
func (m *SQLMap) Delete(key string) error {
	return nil
}

// конец решения

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m, err := NewSQLMap(db)
	if err != nil {
		panic(err)
	}

	m.Set("name", "Alice")
	m.Set("age", 42)

	name, err := m.Get("name")
	fmt.Printf("name = %v, err = %v\n", name, err)
	// name = Alice, err = <nil>

	age, err := m.Get("age")
	fmt.Printf("age = %v, err = %v\n", age, err)
	// age = 42, err = <nil>

	m.Set("name", "Bob")
	name, err = m.Get("name")
	fmt.Printf("name = %v, err = %v\n", name, err)
	// name = Bob, err = <nil>

	m.Delete("name")
	name, err = m.Get("name")
	fmt.Printf("name = %v, err = %v\n", name, err)
	// name = <nil>, err = sql: no rows in result set
}