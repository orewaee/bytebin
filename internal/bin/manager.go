package bin

type Manager interface {
	Add(id string, bytes []byte) error
	RemoveById(id string) error
	GetById(id string) ([]byte, error)
}
