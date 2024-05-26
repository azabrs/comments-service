package storage

type Storage interface{
	Register(string, string) error
	IsRegister(string) (bool, error)
	//ToDo
}