package data

import (
	"be_demo/internal/conf"
	"be_demo/internal/infrastructure/library"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type DBS struct {
	RwDb *gorm.DB
	RoDb *gorm.DB
	Log  *log.Helper
}

func NewMysqlDBS(
	conf *conf.Bootstrap,
	logger log.Logger,
) (*DBS, func(), error) {
	logDb := log.NewHelper(log.With(logger, "x_module", "data/NewMysqlDBS"))
	rwDb := NewDb(conf.GetData().GetRwdb(), logger)
	roDb := NewDb(conf.GetData().GetRodb(), logger)
	clean := func() {
		if rwDb != nil {
			if db, err := rwDb.DB(); err != nil {
				db.Close()
			}
		}
		if roDb != nil {
			if db, err := roDb.DB(); err != nil {
				db.Close()
			}
		}
	}
	return &DBS{
		RwDb: rwDb,
		RoDb: roDb,
		Log:  logDb,
	}, clean, nil
}

// NewDb
func NewDb(conf *conf.Data_Database, lg log.Logger) *gorm.DB {
	logDb := log.NewHelper(log.With(lg, "x_module", "data/NewDb"))
	logDb.Infof("init %s...", conf.GetDriver())

	db, err := gorm.Open(mysql.Open(conf.Source), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 library.NewGorm(lg),
	})
	if err != nil {
		logDb.Fatalf("failed opening connection to mysql: %v", err)
	}

	db = db.Session(&gorm.Session{
		Logger: db.Logger.LogMode(logger.Info),
	})
	err = db.Use(
		dbresolver.Register(dbresolver.Config{Replicas: []gorm.Dialector{mysql.Open(conf.Source)}}).
			SetConnMaxLifetime(time.Hour).
			SetMaxIdleConns(int(conf.MaxIdleConns)).
			SetMaxOpenConns(int(conf.MaxOpenConns)),
	)
	if err != nil {
		logDb.Fatalf("failed dbr use to mysql: %v", err)
	}
	return db
}
