package father


const EMPTY_CACHE_STRING = "nil"
const DRIVE_NAME = "mysql"

const TIMESTAMP6 = "2006-01-02 15:04:05.999999"

const (
	KEYWORD_AUTOINCREMENT     = "_ai"
	KEYWORD_PRIMARY_KEY       = "_pk"
	KEYWORD_UNIQUE_KEY        = "_uq"
	KEYWORD_CREATE_TIME       = "_ct"
	KEYWORD_UPDATE_TIME       = "_ut"
	KEYWORD_CREATE_TIMESTAMP6 = "_cts6"
	KEYWORD_UPDATE_TIMESTAMP6 = "_uts6"
)

const (
	TYPE_SELECT int = iota
	TYPE_INSERT
	TYPE_UPDATE
	TYPE_DELETE
)


const LOGICAL_AND   = "AND"
const LOGICAL_OR    = "OR"
const BRACKET_OPEN  = "("
const BRACKET_CLOSE = ")"
