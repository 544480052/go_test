package father

import (
	"time"
)

type EventParams struct {
	Session      *Session
	Table        string
	Sql          string
	Args         []interface{}
	Params       *map[string]interface{}
	Before       interface{}
	After        interface{}
	Fields       []string
	ColIdx       []int
	InsertId     int64
	RowsAffected int64
	StartTime    time.Time
	QueryTime    float64
	Status       bool
	Uri          *Uri
}

type EventCall func(params *EventParams) error

func (o *Orm) eventQuery(s *Session) (err error) {
	if o.onQuery == nil {
		return
	}

	return o.eventCall(s, o.getEventParams(s), o.onQuery)
}

func (o *Orm) eventUpdate(s *Session, obj interface{}) (err error) {
	ep := o.getEventParams(s)
	ep.After  = obj
	ep.Fields = []string{}

	for _, f := range s.table.Fields {
		if _, ok := s.set[f]; ok {
			ep.Fields = append(ep.Fields, f)
		}
	}

	if dao, ok := obj.(Daoer); ok {
		ep.Before = dao.Before()
		ep.ColIdx = dao.ColIdx()
	}

	return o.eventCall(s, ep, o.onUpdate)
}

func (o *Orm) eventInsert(s *Session, obj interface{}) (err error) {
	ep := o.getEventParams(s)
	ep.After = obj

	if dao, ok := obj.(Daoer); ok {
		ep.Before = dao.Before()
		ep.ColIdx = dao.ColIdx()
	}

	return o.eventCall(s, ep, o.onInsert)
}

func (o *Orm) eventDelete(s *Session) (err error) {
	return o.eventCall(s, o.getEventParams(s), o.onDelete)
}

func (o *Orm) eventLongQuery(s *Session) {
	if o.onLongQuery != nil {
		go o.onLongQuery(o.getEventParams(s))
	} else {
		o.log.Notice(*o.getEventParams(s))
	}
}

func (o *Orm) getEventParams(s *Session) *EventParams {
	return &EventParams{
		Session:      s,
		Table:        s.table.Name,
		Sql:          s.sql,
		Args:         s.args,
		Params:       s.params,
		InsertId:     s.insertId,
		RowsAffected: s.rowsAffected,
		StartTime:    s.queryStart,
		QueryTime:    s.queryTime,
		Status:       s.status,
		Uri:          o.uri,
	}
}

func (o *Orm) eventCall(s *Session, ep *EventParams, fn EventCall) (err error) {
	s.reset()

	if s.tx != nil {
		s.txId ++
		err = fn(ep)
		s.txId --
	} else {
		go fn(ep)
	}

	return
}