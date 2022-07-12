package repository

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"go-http-server/domain"
	"go-http-server/usecase/repository"
)

type indexRepository struct {
	Template *template.Template
}

func NewIndexRepository(t *template.Template) repository.IndexRepository {
	return &indexRepository{
		Template: t,
	}
}

func (ir *indexRepository) SetView(w http.ResponseWriter, r *http.Request) error {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed Get current path. due to an error: %v\n", err)
		return err
	}

	data := domain.IndexPage{
		Title:  "Banner Sample via Golang",
		Header: "Index",
	}

	t, err := ir.Template.ParseFiles(currentPath + "/static/index.html")
	if err != nil {
		fmt.Printf("Failed Create template. due to an error: %v\n", err)
		return err
	}

	if err := t.Execute(w, data); err != nil {
		fmt.Printf("Failed Execute template. due to an error: %v\n", err)
		return err
	}

	return nil
}
