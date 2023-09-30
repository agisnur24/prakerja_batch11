package article

type Article struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserId      uint   `gorm:"foreignKey" json:"user_id"`
}

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
