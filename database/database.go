package database

type DbService interface {
	GetDb() any
	CloseDb()
}

type DbStub interface {
	FetchById(...string) any
}
