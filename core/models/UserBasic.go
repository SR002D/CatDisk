package models

type UserBasic struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

//func (table *UserBasic) Insert(db *sql.DB) error {
//
//}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
