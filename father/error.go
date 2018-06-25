package father

import "errors"

var (
	ErrMissingDao           = errors.New("missing dao")
	ErrUnknownStruct        = errors.New("unknown struct")
	ErrSqlEmpty             = errors.New("sql is empty")
	ErrNeedPointer          = errors.New("need a pointer")
	ErrNeedPtrToSlice       = errors.New("need a pointer to a slice")
	ErrNoPrimaryKey         = errors.New("not primary key")
	ErrNoPrimaryAndUnique   = errors.New("not primary key and unique key")
	ErrUpdateParamsEmpty    = errors.New("update params empty")
	ErrWhereEmpty           = errors.New("where empty")
	ErrDuplicateValues      = errors.New("duplicate values")
	ErrNotSetUpdateField    = errors.New("not set update field")
	ErrTransExist           = errors.New("transaction has exist")
	ErrTransNotExist        = errors.New("transaction not exist")
)

