package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DefaultMaxIdleConnections = 10
	DefaultMaxOpenConnections = 100
)

// New returns a new gorm.DB instance with options.
//
// The default options are:
//  - MaxIdleConnections: DefaultMaxIdleConnections
//  - MaxOpenConnections: DefaultMaxOpenConnections
//  - EnableAutoMigrate: true
//
func New(tables []interface{}, opts ...Option) (*gorm.DB, error) {
	o := Options{
		MaxIdleConnections: DefaultMaxIdleConnections,
		MaxOpenConnections: DefaultMaxOpenConnections,
		EnableAutoMigrate:  true,
	}

	for _, opt := range opts {
		opt(&o)
	}

	db, err := gorm.Open(mysql.Open(o.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if err = configure(db, &o); err != nil {
		return nil, err
	}

	if o.EnableAutoMigrate {
		if err = db.AutoMigrate(tables...); err != nil {
			return nil, err
		}
	}

	return db, err
}

// configure gorm.
func configure(db *gorm.DB, o *Options) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)
	return nil
}
