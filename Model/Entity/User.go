package Entity

type User struct {
	ID       int    `json:"id,omitempty" gorm:"primaryKey"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
