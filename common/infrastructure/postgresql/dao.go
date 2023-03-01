package postgresql

import "errors"

var (
	errRowExists    = errors.New("row exists")
	errRowNotExists = errors.New("row doesn't exist")
)

type dbTable struct {
	name string
}

func NewDBTable(name string) dbTable {
	return dbTable{name: name}
}

func (t dbTable) Insert(filter, result interface{}) error {
	query := db.Table(t.name).Where(filter).FirstOrCreate(result)

	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected == 0 {
		return errRowExists
	}

	return nil
}

func (t dbTable) GetTableRecords(
	filter, result interface{}, limit, offset int,
	column string, desc bool,
) (_ int, err error) {
	var total int64
	query := db.Table(t.name).Where(filter)
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return 0, err
	}

	if len(column) != 0 {
		if desc {
			query.Order(column + " desc")
		} else {
			query.Order(column + " asc")
		}
	}

	if limit > 0 && offset >= 0 {
		query.Limit(limit).Offset(offset)
	}

	err = query.Find(result).Error

	return int(total), err
}

func (t dbTable) IsRowNotExists(err error) bool {
	return errors.Is(err, errRowNotExists)
}

func (t dbTable) IsRowExists(err error) bool {
	return errors.Is(err, errRowExists)
}
