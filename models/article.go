package models

type (
	Article struct {
		Id	uint64 `json:"Id"`
		Title string `json:"title"`
		Desc string `json:"desc"`
		Content string `json:"content"`
		AuthorID uint64 `json:"author"`
	}
)
