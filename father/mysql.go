package father

import (
	"sync"
	"time"
	"runtime"
	"math/rand"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	/**
	import "fmt"	最常用的一种形式
	import "./test"	导入同一目录下test包中的内容
	import f "fmt"	导入fmt，并给他启别名ｆ
	import . "fmt"	将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加fmt，如fmt.Println可以直接写成Println
	import  _ "fmt" 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容
	 */
)

type Orm struct {
	/**
	a 	myStruct    	a是结构体类型变量
	b 	*myStruct		b是指向一个结构体类型变量的指针
	c	[]*myStruct		c是指向一个结构体类型变量的指针数组
	 */


	db         *sql.DB
	dbSlave    []*sql.DB
	dbSlaveLen int

	// 开启 debug 会打印查询日志
	debug bool

	// log handler
	log Logger

	// 缓存
	cache       Cacher
	enableCache bool
	cacheTime   int
	cacheEmpty  bool

	// uri 包含连接相关信息
	// fixme
	uri *Uri

	// 慢查询时间阙值
	longQueryTime float64

	// 事件
	onQuery  EventCall
	onUpdate EventCall
	onInsert EventCall
	onDelete EventCall

	onLongQuery EventCall

	// 表结构缓存
	tableCache sync.Map

	cachePrefix     string
	primaryCacheKey string
	uniqueCacheKey  string
}

func init() {
	// 随机种子
	rand.Seed(time.Now().UnixNano())
}

//  实例化一个 ORM
func Open(dataSource string) (*Orm, error) {
	db, err := sql.Open(DRIVE_NAME, dataSource)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	o := &Orm{
		db:  db,
		log: newLogger(),
	}

	o.uri = o.ParseMySQLUri(DRIVE_NAME, dataSource)

	// 结束的时候执行，关闭数据库连接
	runtime.SetFinalizer(o, func(o *Orm) {
		o.Close()
	})

	return o, nil
}

// 添加 Slave
func (o *Orm) AddSlave(dataSource string) error {
	db, err := sql.Open(DRIVE_NAME, dataSource)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	o.dbSlave = append(o.dbSlave, db)
	o.dbSlaveLen++

	return nil
}

// 设置数据库最大空闲连接数
func (o *Orm) SetMaxIdleConns(n int) *Orm {
	o.db.SetMaxIdleConns(n)

	return o
}

// 设置数据库最大连接数
func (o *Orm) SetMaxOpenConns(n int) *Orm {
	o.db.SetMaxOpenConns(n)

	return o
}

// 设置连接最大存活时间
func (o *Orm) SetConnMaxLifetime(d time.Duration) *Orm {
	o.db.SetConnMaxLifetime(d)

	return o
}

// 设为 Debug Model
func (o *Orm) SetDebug(debug bool) *Orm {
	o.debug = debug

	return o
}

// 设置慢查询阙值
func (o *Orm) SetLongQueryTime(time float64) *Orm {
	o.longQueryTime = time

	return o
}

// 设置是否启用缓存
func (o *Orm) SetEnableCache(cache bool, prefix ...string) *Orm {
	o.enableCache = cache

	if cache && o.cachePrefix == "" {
		if len(prefix) > 0 {
			o.cachePrefix = prefix[0]
		} else {
			o.cachePrefix = "lemon"
		}

		o.primaryCacheKey = o.cachePrefix + ":primary:%s:%s"
		o.uniqueCacheKey = o.cachePrefix + ":unique:%s:%s"
	}

	return o
}

// 设置缓存时间
func (o *Orm) SetCacheTime(second int) *Orm {
	o.cacheTime = second

	return o
}

func (o *Orm) SetLogHandler(log Logger) *Orm {
	o.log = log

	return o
}

// 设置缓存器
func (o *Orm) SetCacheHandler(ch Cacher) *Orm {
	o.cache = ch

	return o
}

// 设置缓存空值
func (o *Orm) SetCacheEmpty(cache bool) *Orm {
	o.cacheEmpty = cache

	return o
}

// 查询事件
func (o *Orm) OnQuery(fn EventCall) *Orm {
	o.onQuery = fn

	return o
}

// 更新事件
func (o *Orm) OnUpdate(fn EventCall) *Orm {
	o.onUpdate = fn

	return o
}

// 插入事件
func (o *Orm) OnInsert(fn EventCall) *Orm {
	o.onInsert = fn

	return o
}

// 删除事件
func (o *Orm) OnDelete(fn EventCall) *Orm {
	o.onDelete = fn

	return o
}

// 慢查询事件
func (o *Orm) OnLongQuery(fn EventCall) *Orm {
	o.onLongQuery = fn

	return o
}

// 连接状态
func (o *Orm) Stats() sql.DBStats {
	return o.db.Stats()
}

// 手动关闭数据库连接
func (o *Orm) Close() error {
	return o.db.Close()
}

// 开启一个新的查询
func (o *Orm) NewSession() *Session {
	return &Session{orm: o, enableCache: o.enableCache}
}

// 随机获取一个从库连接
func (o *Orm) getSlave() *sql.DB {
	return o.dbSlave[rand.Intn(o.dbSlaveLen)]
}
