package todo

type Todo struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"Not Null" json:"name"`
}
