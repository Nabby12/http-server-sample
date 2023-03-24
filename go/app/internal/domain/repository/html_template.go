package repository

import (
	"go-http-server/internal/domain/model"
)

type HtmlTemplate interface {
	GetByName(key string) (model.HtmlTemplate, error)
}
