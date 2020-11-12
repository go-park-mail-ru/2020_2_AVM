package models

type (
	Article struct {
		Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		ArticleTitle string `json:"article_title" gorm:"column:article_title"`
		Description string `json:"description" gorm:"column:description"`
		Content string `json:"content" gorm:"column:content"`
		CategoryID uint64 `json:"categoryid" gorm:"column:categoryid"` //Foreign key
		AuthorID uint64 `json:"authorid" gorm:"column:authorid"`		//Foreign key
	}

	Tag struct {
		Id	uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		TagTitle string `json:"tag_title" gorm:"column:tag_title"`
	}
	Tags struct {
		TagsValues []Tag `json:"tags"`
	}

	TagArticle struct {
		Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		ArticleID uint64 `json:"articleid" gorm:"column:articleid"`
		TagID uint64 `json:"tagid" gorm:"column:tagid"`
	}


	Category struct {
		Id	uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		CategoryTitle string `json:"category_title" gorm:"column:category_title"`
	}

	Categories struct {
		CategoriesValues []Tag `json:"categories"`
	}
)
