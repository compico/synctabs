package booklink

type DataInterface interface {
	FromJson() error
	ToJson() ([]byte, error)
}

func ToJson(di DataInterface) ([]byte, error) {
	return di.ToJson()
}