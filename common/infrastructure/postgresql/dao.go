package postgresql

import (
	"errors"
)

var (
	errRowExists    = errors.New("row exists")
	errRowNotExists = errors.New("row doesn't exist")
)

type SortByColumn struct {
	Column string
	Ascend bool
}

func (s SortByColumn) order() string {
	v := " ASC,"
	if !s.Ascend {
		v = " DESC,"
	}
	return s.Column + v
}

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
	filter, result interface{}, pageNum, countPerPage int,
	sort []SortByColumn,
) (err error) {
	query := db.Table(t.name).Where(filter)

	var order string
	for _, v := range sort {
		order += v.order()
	}

	if len(order) >= 0 {
		query.Order(order[:len(order)-1])
	}

	if countPerPage > 0 {
		offset := 0
		if pageNum > 0 {
			offset = (pageNum - 1) * countPerPage
		}
		query.Limit(countPerPage).Offset(offset)
	}

	err = query.Find(result).Error

	return
}

func (t dbTable) Counts(filter interface{}) (int, error) {
	var total int64
	err := db.Table(t.name).Where(filter).Count(&total).Error

	return int(total), err
}

func (t dbTable) IsRowNotExists(err error) bool {
	return errors.Is(err, errRowNotExists)
}

func (t dbTable) IsRowExists(err error) bool {
	return errors.Is(err, errRowExists)
}
