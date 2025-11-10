package database

import (
	"fmt"
	"runtime"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gmLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var gormDB map[string]*gorm.DB

func InitDB(c []*MySqlConfig) error {
	return initMysql(c)
}

func initMysql(c []*MySqlConfig) error {
	gormDB = make(map[string]*gorm.DB, len(c))
	for _, m := range c {
		checkDbConfig(m)
		masterDSN := genDbDSN(m)

		gmDB, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       masterDSN,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: gmLogger.Default.LogMode(gmLogger.Info),
		})
		if err != nil {
			return fmt.Errorf("mysql connect %v", err)
		}

		// 注册读写分离
		slaves := make([]gorm.Dialector, 0, len(m.Slave))
		for _, s := range m.Slave {
			checkDbConfig(s)
			slaves = append(slaves, mysql.Open(genDbDSN(s)))
		}
		if err = gmDB.Use(dbresolver.Register(dbresolver.Config{
			Sources:           []gorm.Dialector{mysql.Open(masterDSN)},
			Replicas:          slaves,
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		}).SetMaxIdleConns(m.MaxIdleConns).SetMaxOpenConns(m.MaxOpenConns).
			SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime)).
			SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifetime))); err != nil {
			return fmt.Errorf("gmDB.Use error: %+v", err)
		}

		// ping
		if db, err := gmDB.DB(); err != nil {
			return fmt.Errorf("mysql sql.DB error role %s: %w", m.Role, err)
		} else {
			if err := db.Ping(); err != nil {
				return fmt.Errorf("mysql ping error role %s: %w", m.Role, err)
			}
		}

		//save gorm
		gormDB[m.Role] = gmDB
	}

	return nil
}

func CloseDB() {
	//保证传入的参数在这个方法被调用之前不被垃圾回收器回收掉
	runtime.KeepAlive(gormDB)
	for k, v := range gormDB {
		if db, err := v.DB(); err == nil {
			db.Close()
		}
		v = nil
		delete(gormDB, k)
	}
	gormDB = nil
}

func GetDB(role ...string) *gorm.DB {
	if len(role) == 0 {
		return gormDB[RoleDefault]
	}
	return gormDB[role[0]]
}

// 处理DB配置
func checkDbConfig(c *MySqlConfig) {
	if c.Charset == "" {
		c.Charset = DbCharset
	}
	if c.Collation == "" {
		c.Collation = DbCollation
	}
	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = DbMaxOpenConns
	}
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = DbMaxIdleConns
	}
	if c.MaxLifetime == 0 {
		c.MaxLifetime = DbMaxLifetime
	}
	if c.MaxIdleTime == 0 {
		c.MaxIdleTime = DbMaxIdleTime
	}
}

// DSN
func genDbDSN(m *MySqlConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s", m.Username, m.Password, m.Host, m.Port, m.Dbname, m.Charset, m.Collation)
}
