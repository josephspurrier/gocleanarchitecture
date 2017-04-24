package jsonrepo

// IRepoService is the interface for storage.
type IRepoService interface {
	Records(recordType string) ([]interface{}, error)
	AddRecord(recordType string, record interface{}) error
}
