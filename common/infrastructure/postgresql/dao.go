package postgresql

type dbCollection struct {
	name string
}

func NewDBCollection(name string) dbCollection {
	return dbCollection{name: name}
}

func (d dbCollection) NewRecordIfNotExists(filter, result interface{}) (rows int64, err error) {
	query := db.Table(d.name).Where(filter).FirstOrCreate(result)

	err = query.Error
	rows = query.RowsAffected

	return
}
