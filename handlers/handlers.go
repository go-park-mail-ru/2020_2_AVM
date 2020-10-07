package handlers

import "github.com/go-park-mail-ru/2020_2_AVM/models"

type (
	Handler struct {
		Articles []models.Article
		Profiles []models.Profile
		logInIds map[string]string
		user_id int 
		article_id int 
	}
)
func NewHandler() (*Handler) {
	return &Handler{nil, nil, map[string]string{}, 0, 0}
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

