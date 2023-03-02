package postgresql

import (
	"errors"
	"fmt"
)

var (
	errRowExists    = errors.New("row exists")
	errRowNotExists = errors.New("row doesn't exist")
	Descend         = Sort("DESC")
	Ascend          = Sort("ASC")
)

type Sort string

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
	sort map[string]Sort,
) (err error) {
	query := db.Table(t.name).Where(filter)

	var order string
	for k, v := range sort {
		if v == Descend {
			order += fmt.Sprintf("%s %s ,", k, Descend)
		}

		if v == Ascend {
			order += fmt.Sprintf("%s %s ,", k, Ascend)
		}
	}

	if len(order) >= 0 {
		query.Order(order[:len(order)-1])
	}

	if pageNum > 0 {
		query.Limit(countPerPage).Offset((pageNum - 1) * countPerPage)
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
