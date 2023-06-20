package Entity

type User struct {
	ID       int    `json:"id,omitempty" gorm:"primaryKey"`
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
