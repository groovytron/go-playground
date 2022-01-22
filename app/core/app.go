package core

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router   *mux.Router
	Database *gorm.DB
}
