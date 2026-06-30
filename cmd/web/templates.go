package main

import (
	"snippetbox.igorcrispim.net/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
