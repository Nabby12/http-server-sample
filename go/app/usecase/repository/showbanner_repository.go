package repository

import (
	"net/http"
	"time"
)

type ShowbannerRepository interface {
	SetView(http.ResponseWriter, string, time.Time, time.Time, time.Time) error
}
