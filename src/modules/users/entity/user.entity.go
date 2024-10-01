package entity

type User struct {
	ID    uint   `json:"id"         gorm:"<-:create"`
	Name  string `json:"name"       gorm:"<-:update"`
	Email string `json:"email"      gorm:"<-:update"`
}
