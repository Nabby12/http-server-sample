package repository

import (
	"go-http-server/internal/domain/model"
	"net/http"
)

type ShowBannerPage interface {
	SetView(model.ShowBannerPage, http.ResponseWriter) error
}
