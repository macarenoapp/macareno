package profile

type Profile struct {
	Id	    	uint	 `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Name        string   `gorm:"column:name"`
	Surname		string   `gorm:"column:surname"`
	Description string   `gorm:"column:description"`
	Bio         string   `gorm:"column:Bio"`
	User	    uint	 `gorm:"column:user"`
}

type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;"`
	Bio          string  `gorm:"column:bio;"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;"`
	Type		 string	 `gorm:"column:type;"` 
}


