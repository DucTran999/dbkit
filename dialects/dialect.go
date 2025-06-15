package dialects

import "gorm.io/gorm"

type Dialect interface {
	Open() (*gorm.DB, error)
}
