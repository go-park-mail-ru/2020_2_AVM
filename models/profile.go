package models

type (
	Profile struct {
		Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		Login string `json:"login" gorm:"column:login"`
		Email string `json:"email" gorm:"column:email"`
		Name string `json:"name" gorm:"column:name"`
		Surname string `json:"surname" gorm:"column:surname"`
		Password string `json:"password" gorm:"column:password"`
		Avatar string `json:"avatar" gorm:"column:avatar"`
		Cookie string `json:"-" gorm:"column:cookie"`
	}

	CategoryFollow struct {
		Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
		ProfileID uint64 `json:"profileid" gorm:"column:profileid"`
		CategoryID uint64 `json:"categoryid" gorm:"column:categoryid"`
	}
)