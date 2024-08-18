package repos

type BinRepo interface {
	AddBin(id string, bin []byte) error
	RemoveBinById(id string) error
	GetBinById(id string) ([]byte, error)
	GetAllBinIds() ([]string, error)
}
