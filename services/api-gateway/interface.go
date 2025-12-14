package gateway

type DataProvider interface {
	fetchData() error
}
