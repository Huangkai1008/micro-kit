package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	metav1 "github.com/Huangkai1008/micro-kit/pkg/meta/v1"
)

type GormRepository struct {
	// gorm db connection.
	Db *gorm.DB

	// logger for gorm.
	logger *zap.Logger // repo logger
}

func NewGormRepository(db *gorm.DB, logger *zap.Logger) *GormRepository {
	return &GormRepository{
		Db:     db,
		logger: logger,
	}
}

func (r *GormRepository) Get(id int) (record metav1.Resource, err error) {
	err = r.Db.Limit(1).Find(&record, id).Error
	return
}

func (r *GormRepository) Find(conditions interface{}) (record metav1.Resource, err error) {
	err = r.Db.Where(conditions).Limit(1).Find(&record).Error
	return
}

func (r *GormRepository) FindAll(conditions interface{}) (records []metav1.Resource, err error) {
	err = r.Db.Where(conditions).Find(&records).Error
	return
}

func (r *GormRepository) Exist(conditions interface{}) (bool, error) {
	var (
		err    error
		record metav1.ObjectMeta
	)

	err = r.Db.Where(conditions).Limit(1).Find(&record).Error
	if err != nil {
		return false, err
	} else {
		return record == metav1.ObjectMeta{}, err
	}
}

func (r *GormRepository) Create(record metav1.Resource) error {
	err := r.Db.Create(record).Error
	return err
}
