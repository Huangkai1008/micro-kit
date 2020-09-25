package gorm

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/Huangkai1008/micro-kit/pkg/message"
)

// New returns a new gorm.DB instance with options.
func New(o *Options, tables []interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.DSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, message.DatabaseConnectError)
	}

	if err = configure(db, o); err != nil {
		return nil, errors.WithMessage(err, message.ORMConfigError)
	}

	if o.EnableAutoMigrate {
		if err = db.AutoMigrate(tables...); err != nil {
			return nil, errors.Wrap(err, message.DatabaseMigrateError)
		}
	}

	return db, err
}

// configure gorm.
func configure(db *gorm.DB, o *Options) error {
	sqlDB, err := db.DB()
	if err != nil {
		return errors.Wrap(err, message.GetConnectionError)
	}
	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)
	return nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
