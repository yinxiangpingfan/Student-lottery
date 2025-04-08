package modles

type Student struct {
	Name string `gorm:"name" json:"name"`
	ID   string `gorm:"id" json:"id"`
	Sex  string `gorm:"sex" json:"sex"`
}
