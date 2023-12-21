package global

import "github.com/jinzhu/gorm"

// DBEngine is a global variable that represents the database engine used in the project.
var (
	DBEngine *gorm.DB
)
