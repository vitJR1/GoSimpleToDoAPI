package entity

type Todo struct {
	ID          uint   `json:"id"          gorm:"<-:create"`
	Title       string `json:"title"       gorm:"<-:update"`
	Description string `json:"description" gorm:"<-:update"`
}
