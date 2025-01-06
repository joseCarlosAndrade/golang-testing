package service

type DataProvider interface {
	GetRandomData(int) (int, error)
}

