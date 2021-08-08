package server

import (
	"net/http"

	"gorm.io/gorm"
)

type Server struct {
	DB         *gorm.DB
	HTTPServer *http.Server
}
