package repo

import (
    "context"
    "database/sql"
    "fmt"
    "github.com/jinzhu/gorm"
    "log"
    "time"
    "unity.service/apps/v1/dto/common"
)

var (
    ServiceDBA *Dba
    err        error
)

type (
    DataBase struct {
        Type        string
        User        string
        Password    string
        Host        string
        DBName      string
        TablePrefix string
        PortTcp     string
        PortHttp    string
        Debug       string
    }

    Dba struct {
        Mysql      *gorm.DB
        ClickHouse *sql.DB
    }
)

func init() {
    ServiceDBA = &Dba{}
}

func Registry(database DataBase) {

    switch database.Type {
    case `mysql`:
        ServiceDBA._mysql_conn(database)
    case `clickhouse`:
        ServiceDBA._click_house_conn(database)
    default:

    }

}

/**
 * _click_house_conn
 * @Description:
 * @receiver d
 * @param database
 * @return *sql.DB
 */
func (d *Dba) _click_house_conn(database DataBase) *sql.DB {

   /* d.ClickHouse, _ = sql.Open(database.Type, fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s&debug=true",
        database.Host, database.PortTcp, database.User, database.Password, database.DBName))
    if err = d.ClickHouse.Ping(); err != nil {
        if exception, ok := err.(*clickhouse.Exception); ok {
            fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
        } else {
            panic(err)
        }
    }
    return d.ClickHouse
*/
    return nil
}

/**
 * _mysql_conn
 * @Description:
 * @receiver d
 * @param database
 * @return *gorm.DB
 */
func (d *Dba) _mysql_conn(database DataBase) *gorm.DB {

    d.Mysql = connection(database).LogMode(false).Debug()
    if err = d.Mysql.DB().Ping(); err != nil {
        _ = d.Mysql.Close()
        d.Mysql = d.Mysql.New()
    }
    return d.Mysql

}

func connection(setting DataBase) *gorm.DB {
    db, err := gorm.Open(setting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        setting.User,
        setting.Password,
        setting.Host,
        setting.DBName))
    if err != nil {
        log.Fatalln("models.Setup err:", err)
    }
    db.SingularTable(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
    db.DB().SetConnMaxLifetime(time.Hour)
    return db
}


// PaginationResult ??????????????????
type PaginationResult struct {
    Total    int  `json:"total"`     // ??????
    Current  uint `json:"current"`   // ?????????
    PageSize uint `json:"page_size"` // ?????????
}

// PaginationParam ????????????
type PaginationParam struct {
    Pagination bool `json:"pagination" form:"pagination"`                            // ????????????????????????
    OnlyCount  bool `json:"-" form:"-"`                                              // ???????????????count
    Current    uint `json:"current" form:"current,default=1"`                        // ?????????
    PageSize   uint `json:"page_size" form:"page_size,default=10" binding:"max=100"` // ?????????
}

// WrapPageQuery ??????????????????????????????
func WrapPageQuery(db *gorm.DB, pp common.PaginationParam, out interface{}) (*common.PaginationResult, error) {
    if pp.OnlyCount {
        var count int64
        err := db.Count(&count).Error
        if err != nil {
            return nil, err
        }
        return &common.PaginationResult{Total: int(count)}, nil
    } else if !pp.Pagination {
        err := db.Find(out).Error
        return nil, err
    }
    total, err := findPage(db, pp, out)
    if err != nil {
        return nil, err
    }
    return &common.PaginationResult{
        Total:    total,
        Current:  pp.Current,
        PageSize: pp.PageSize,
    }, nil
}

func findPage(db *gorm.DB, pp common.PaginationParam, out interface{}) (int, error) {
    var count int64
    err := db.Count(&count).Error
    if err != nil {
        return 0, err
    } else if count == 0 {
        return int(count), nil
    }
    current, pageSize := int(pp.Current), int(pp.PageSize)
    if current > 0 && pageSize > 0 {
        db = db.Offset((current - 1) * pageSize).Limit(pageSize)
    } else if pageSize > 0 {
        db = db.Limit(pageSize)
    }

    err = db.Find(out).Error
    return int(count), err
}

// TransRepo ??????
type TransRepo struct {
    DB *gorm.DB
}

type (
    transCtx struct{}
)

// NewTrans ??????????????????????????????
func NewTrans(ctx context.Context, db *gorm.DB) context.Context {
    return context.WithValue(ctx, transCtx{}, db)
}

// FromTrans ????????????????????????
func FromTrans(ctx context.Context) (interface{}, bool) {
    v := ctx.Value(transCtx{})
    return v, v != nil
}

// Exec ????????????
func (a *TransRepo) Exec(ctx context.Context, fn func(context.Context) error) error {
    if _, ok := FromTrans(ctx); ok {
        return fn(ctx)
    }

    return a.DB.Transaction(func(db *gorm.DB) error {
        return fn(NewTrans(ctx, db))
    })
}

func getDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
    trans, ok := FromTrans(ctx)
    if ok {
        db, ok := trans.(*gorm.DB)
        if ok {
            return db
        }
    }

    return defDB
}

// Tabler ????????????
type Tabler interface {
    TableName() string
}

// GetDBWithTable ????????????struct?????????db
func GetDBWithTable(ctx context.Context, taber Tabler) *gorm.DB {
    return getDB(ctx, ServiceDBA.Mysql).Table(taber.TableName())
}

// NewTransRepo ??????repo
func NewTransRepo() *TransRepo {
    return &TransRepo{
        DB: ServiceDBA.Mysql,
    }
}
