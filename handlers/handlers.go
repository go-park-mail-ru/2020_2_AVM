package handlers

import (
	"github.com/go-park-mail-ru/2020_2_AVM/models"
)

type (
	Handler struct {
		Articles []models.Article
		Profiles []models.Profile
		logInIds map[string]string
		userId int
		articleId int
	}
)
	
func (h *Handler) GetNewUserId() (int) {
	h.userId += 1
	return h.userId
}

func (h *Handler) GetNewArcticleId() (int) {
	h.articleId += 1
	return h.articleId
}

func NewHandler() (*Handler) {
	return &Handler{nil, nil, map[string]string{}, 0, 0}
}