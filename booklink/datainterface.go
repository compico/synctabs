package booklink

type DataInterface interface {
	ToJson() ([]byte, error)
}

func ToJson(data DataInterface) ([]byte, error) {
	return data.ToJson()
}
