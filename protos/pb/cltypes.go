package pb

import (
	sq "github.com/Masterminds/squirrel"
)

func (x *Int64Comp) Where(column string, rq sq.SelectBuilder) sq.SelectBuilder {
	if x == nil {
		return rq
	}
	switch x.Op {
	case ">", "<", ">=", "<=", "=", "!=":
		return rq.Where(column+" "+x.Op+" ?", x.Value)
	}
	return rq.Where(column+" = ?", x.Value)
}
