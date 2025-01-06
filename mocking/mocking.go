package mocking

type DataBaseInterface interface {
	GetUser(id int) (string, error)
}