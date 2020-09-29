package repository

import metav1 "github.com/Huangkai1008/micro-kit/pkg/meta/v1"

// Repository is an interface for a repository.
type Repository interface {
	// Get returns the record for the given id.
	Get(id int) (record metav1.Resource, err error)
	// Find return one record filter by the conditions.
	Find(conditions interface{}) (record metav1.Resource, err error)
	// FindAll return all records filter by the conditions.
	FindAll(conditions interface{}) (records []metav1.Resource, err error)

	// Exist return one record does exist in table.
	Exist(conditions interface{}) (bool, error)

	// Create record.
	Create(record metav1.Resource) error
}
