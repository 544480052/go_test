package father

import (
	"database/sql"
	"fmt"
	"time"
)

type conditionStore struct {
	column    string
	value     interface{}
	operator  string
	connector string
	bracket   string
}

type rawStore struct {
	value string
}

type Session struct {
	orm *Orm
	tx  *sql.Tx

	txId int

	useSlave    int // 0: auto, 1: slave, 2: master
	enableCache bool

	status bool
	error  error

	sql  string
	args []interface{}

	insertId     int64
	rowsAffected int64

	queryStart time.Time
	queryTime  float64

	options []string
	columns []string
	orderBy []string
	groupBy []string
	where   []conditionStore
	table   *TableInfo
	limit   int
	offset  int

	set           map[string]interface{}  // for update
	fields        []string                // for insert
	colIdx        []int                   // for select
	params        *map[string]interface{} // for insert、update cb
	onCommitTasks []func()                // after transaction commit task
}

// 强制使用 master
func (s *Session) Master() *Session {
	s.useSlave = 2

	return s
}

// 强制使用 slave
func (s *Session) Slave() *Session {
	if s.orm.dbSlaveLen > 0 {
		s.useSlave = 1
	}

	return s
}

// 设置是否启用缓存
func (s *Session) EnableCache(cache bool) *Session {
	s.enableCache = cache

	return s
}

// 查询后执行相关操作，各种事件回掉，打印 SQL 等
func (s *Session) after(t string, status bool) {
	s.queryTime = time.Since(s.queryStart).Seconds()
	s.status = status

	if s.orm.longQueryTime > 0 && s.queryTime >= s.orm.longQueryTime {
		s.orm.eventLongQuery(s)
	}

	if s.orm.debug {
		str := "%s %v [%fs] \033[49;" + _if(status, "32;1m√", "31;1mx").(string) + "\033[0m"
		s.orm.log.Debug(fmt.Sprintf(str, s.sql, s.args, s.queryTime))
	}
}

// 重设清理
func (s *Session) reset() {
	s.error = nil

	if s.tx == nil {
		s.insertId = 0
		s.rowsAffected = 0
	}

	s.sql = ""

	s.limit = 0
	s.offset = 0
	s.table = nil
	s.options = nil
	s.columns = nil
	s.orderBy = nil
	s.groupBy = nil
	s.where = nil
	s.set = nil
	s.fields = nil
	s.colIdx = nil
	s.args = nil
	s.params = nil
}