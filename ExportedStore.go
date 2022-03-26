package SpindlyExports

type ExportedStore struct {
	Name string
	Get  func() string
	Set  func(string)
	Next func() string
}
