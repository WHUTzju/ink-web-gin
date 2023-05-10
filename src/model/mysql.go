package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	m "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"ink-web/src/config"
	"ink-web/src/global"
	"ink-web/src/util/log"
	"time"
)

type Options struct {
	ctx         context.Context
	driver      string
	uri         string
	tablePrefix string `mapstructure:"table_prefix" json:"tablePrefix"`
}

func InitDataBases(ctx context.Context) (err error) {

	ops := &Options{
		driver:      "mysql",
		uri:         config.DatabaseSetting.Uri,
		tablePrefix: config.DatabaseSetting.TablePrefix,
	}
	_, err = sql.Open(ops.driver, ops.uri)
	if err != nil {
		log.WithContext(ops.ctx).WithError(err).Error("open %s(%s) failed", ops.driver, ops.uri)
		return
	}

	//开始初始化链接
	init := false
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(config.ServerSetting.ReadTimeout)*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				if !init {
					panic(fmt.Sprintf("initialize mysql failed: connect timeout(%ds)", config.ServerSetting.ReadTimeout))
				}
				// avoid goroutine dead lock
				return
			}
		}
	}()
	// 默认设置
	//type Config struct {
	//	SkipDefaultTransaction   bool 跳过默认事务
	//	NamingStrategy           schema.Namer 命名策略
	/*	```
		type Namer interface {
		    TableName(table string) string
		    SchemaName(table string) string
		    ColumnName(table, column string) string
		    JoinTableName(table string) string
		    RelationshipFKName(Relationship) string
		    CheckerName(table, column string) string
		    IndexName(table, column string) string
		}
		```*/

	//	Logger                   logger.Interface 允许通过覆盖此选项更改 GORM 的默认 logger
	//	NowFunc                  func() time.Time 更改创建时间使用的函数
	//	DryRun                   bool
	//	PrepareStmt              bool
	//	DisableNestedTransaction bool
	//	AllowGlobalUpdate        bool
	//	DisableAutomaticPing     bool
	//	DisableForeignKeyConstraintWhenMigrating bool
	//}

	var db *gorm.DB
	db, err = gorm.Open(m.New(m.Config{
		DSN:                       ops.uri,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   false, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   ops.tablePrefix,
				SingularTable: true,
			},
			// select * from xxx => select a,b,c from xxx
			QueryFields:          true,
			DisableAutomaticPing: false,
		})
	if err != nil {
		panic(errors.Wrap(err, "initialize mysql failed"))
	}

	//db.LogMode(viper.GetBool("db.log")) // open debug
	//db.SingularTable(true)
	//db.DB().SetMaxIdleConns(maxIdleConns)
	//db.DB().SetMaxOpenConns(maxOpenConns)

	init = true
	global.Mysql = db

	migrate()
	return nil
}

func migrate() {
	_ = global.Mysql.AutoMigrate(&User{})
}
