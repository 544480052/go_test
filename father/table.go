package father


import (
	"reflect"
	"strings"
)

type TableInfo struct {
	Name          string
	Fields        []string
	PrimaryKeys   []*string
	UniqueKeys    map[string][]*string
	AutoIncrement *string
	Created       *string
	TypeCreated   string
	Updated       *string
	TypeUpdated   string
}

// 从结构体中获取表信息
func (o *Orm) GetTableInfo(table interface{}, isType ...bool) (*TableInfo, error) {
	var v reflect.Value
	var t reflect.Type

	if len(isType) > 0 && isType[0] == true {
		t = table.(reflect.Type)
	} else {
		v = reflect.ValueOf(table)

		if v.Kind() == reflect.Ptr {
			t = v.Elem().Type()
		} else {
			t = v.Type()
		}
	}

	// 结构体名
	sName := t.Name()

	if !v.IsValid() {
		v = reflect.New(t)
	}

	// 是否继承 Dao
	if _, ok := v.Interface().(Daoer); !ok {
		return nil, ErrUnknownStruct
	}

	// 第一个字段是否 Dao
	if _, ok := v.Elem().Field(0).Interface().(Dao); !ok {
		return nil, ErrUnknownStruct
	}

	//  表名
	tName := ""

	// 有 tableName 方法，则从 tableName 方法取得表名
	if fn, ok := v.Interface().(TableName); ok {
		tName = fn.TableName()
	} else {
		// 否则认为结构体名首字母小写为表名
		tName = lcFirst(sName)
	}

	// 是否已经缓存过
	if s, ok := o.tableCache.Load(tName); ok {
		return s.(*TableInfo), nil
	}

	// 如果是指针则取得具体元素
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 缓存表信息
	sc, err := o.cacheTableInfo(t, tName)
	if err != nil {
		return nil, err
	}

	return sc, nil
}

// 缓存表信息
func (o *Orm) cacheTableInfo(t reflect.Type, tName string) (*TableInfo, error) {
	// 表字段数
	fNum := t.NumField()

	// 实例化一个 TableInfo
	newVal := &TableInfo{
		Name:        tName,
		Fields:      make([]string, fNum - 1),
		PrimaryKeys: []*string{},
		UniqueKeys:  map[string][]*string{},
	}

	for i := 1; i < fNum; i++ {
		// 在没有 Tag 的情况下，则字段名为首字母小写的字段名
		newVal.Fields[i - 1] = lcFirst(t.Field(i).Name)

		// 取得 Tag
		tag := t.Field(i).Tag.Get("db")
		if tag == "" {
			continue
		}

		// 指针
		fPtr := &newVal.Fields[i - 1]

		// 去空格，并且按空格分割成 []string
		tags := strings.Fields(tag)

		// 支持无序的定义 eg:`db:customFieldName _cts6` `db:_cts6 customFieldName`
		for _, v := range tags {
			if o.isKeyword(v) {
				switch v {
				case KEYWORD_AUTOINCREMENT:     // 自增长字段
					newVal.AutoIncrement = fPtr
					continue
				case KEYWORD_PRIMARY_KEY:       // 主键字段
					newVal.PrimaryKeys = append(newVal.PrimaryKeys, fPtr)
					continue
				case KEYWORD_CREATE_TIME:       // created int64 类型
				case KEYWORD_CREATE_TIMESTAMP6: // created timestamp 类型
					newVal.Created = fPtr
					newVal.TypeCreated = v
					continue
				case KEYWORD_UPDATE_TIME:       // updated int64 类型
				case KEYWORD_UPDATE_TIMESTAMP6: // updated timestamp 类型
					newVal.Updated = fPtr
					newVal.TypeUpdated = v
					continue
				}

				// 唯一字段
				if v[:3] == KEYWORD_UNIQUE_KEY {
					if u := strings.Replace(v, KEYWORD_UNIQUE_KEY, "", -1); u == "" {
						newVal.UniqueKeys[*fPtr] = append(newVal.UniqueKeys[*fPtr], fPtr)
					} else {
						u := u[1:]
						newVal.UniqueKeys[u] = append(newVal.UniqueKeys[u], fPtr)
					}
				}
			} else {
				// 自定义字段名 eg:`db:"customFieldName"`
				newVal.Fields[i - 1] = v
			}
		}
	}

	// 没有主键
	if len(newVal.PrimaryKeys) == 0 {
		return nil, ErrNoPrimaryKey
	}

	// 存在 orm 实例中
	o.tableCache.Store(tName, newVal)

	return newVal, nil
}

// 是否为索引别名
func (o *Orm) isKeyword(tag string) bool {
	return len(tag) > 0 && tag[:1] == "_"
}