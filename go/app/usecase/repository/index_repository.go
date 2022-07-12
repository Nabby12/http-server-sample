package repository

import (
	"net/http"
)

type IndexRepository interface {
	SetView(http.ResponseWriter, *http.Request) error
}
