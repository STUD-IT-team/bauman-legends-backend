package domain

// ID
//
// Тип первичного ключа (uuid) в таблицах
type ID string

// String
//
// Явное преобразование типа ID в тип string
func (id ID) String() string {
	return string(id)
}
