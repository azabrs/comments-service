package storage

type Storage interface{
	Register(string) error
	IsRegister(string) (bool, error)
	//ToDo
}