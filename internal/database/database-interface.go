package database

type IDatabase interface {
	connect()
	GetClient()
}
