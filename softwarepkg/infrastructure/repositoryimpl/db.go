package repositoryimpl

type dbClient interface {
	Insert(filter, result interface{}) error
	GetTableRecords(filter, result interface{}, limit, offset int, column string, desc bool) (total int, err error)

	IsRowNotExists(err error) bool
	IsRowExists(err error) bool
}
