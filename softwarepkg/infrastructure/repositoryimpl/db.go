package repositoryimpl

type dbClient interface {
	NewRecordIfNotExists(filter, result interface{}) (rows int64, err error)
}
