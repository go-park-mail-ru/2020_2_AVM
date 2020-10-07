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

func (h *Handler) GetNewUserId() (int) {
	h.user_id += 1
	return h.user_id
}

func (h *Handler) GetNewArcticleId() (int) {
	h.article_id += 1
	return h.article_id
}

func NewHandler() (*Handler) {
	return &Handler{nil, nil, map[string]string{}, 0, 0}
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

