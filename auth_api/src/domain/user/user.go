package domain

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"` // -を指定すると非表示にできる
}