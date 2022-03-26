package SpindlyExports

type ExportedStore struct {
	Name string
	Get  func() string
	Set  func(string)
	Next func() string
}

func (es *ExportedStore) GetValue() string {
	return es.Get()
}

func (es *ExportedStore) SetValue(s string) {
	es.Set(s)
}

func (es *ExportedStore) NextValue() string {
	return es.Next()
}
