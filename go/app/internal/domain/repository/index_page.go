package repository

import (
	"go-http-server/internal/domain/model"
	"net/http"
)

type IndexPage interface {
	SetView(model.IndexPage, http.ResponseWriter) error
}
