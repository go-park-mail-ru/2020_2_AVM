package models

type (
	Article struct {
		Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		Title string `json:"title" gorm:"column:title"`
		Desc string `json:"desc" gorm:"column:desc"`
		Content string `json:"content" gorm:"column:content"`
		CategoryID uint64 `json:"categoryid" gorm:"column:categoryid"` //Foreign key
		AuthorID uint64 `json:"authorid" gorm:"column:authorid"`		//Foreign key
	}

	Tag struct {
		Id	uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		Title string `json:"title" gorm:"column:title"`
	}
	TagArticle struct {
		Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		ArticleID uint64 `json:"articleid" gorm:"column:articleid"`
		TagID uint64 `json:"tagid" gorm:"column:tagid"`
	}


	Category struct {
		Id	uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		Title string `json:"title" gorm:"column:title"`
	}

)
