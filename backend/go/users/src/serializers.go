package users

import (
	"github.com/gin-gonic/gin"

	"goApp/common"
)

type ProfileSerializer struct {
	C *gin.Context
	User
}

// Declare your response schema here
type ProfileResponse struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	// Bearer    string  `json:"Bearer"`
	Type 	  string  `json:"type"`
}

// Put your response logic including wrap the userModel here. 
func (self *ProfileSerializer) Response() ProfileResponse {
	// myUserModel := self.C.MustGet("my_user_model").(User)
	profile := ProfileResponse{
		ID:        self.ID,
		Username:  self.Username,
		Email:     self.Email,
		Bio:       self.Bio,
		Image:     self.Image,
		Type:      self.Type,
		// Bearer:    common.GenToken(myUserModel.ID),
	}
	return profile
}

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Id       uint    `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Bearer   string  `json:"Bearer"`
	Type 	 string  `json:"type"`
}

func (self *UserSerializer) Response() UserResponse {
	myUserModel := self.c.MustGet("my_user_model").(User)
	user := UserResponse{
		Id:		  myUserModel.ID,
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Bearer:   common.GenToken(myUserModel.ID),  //Aqui crea el token en una funcion de common/utils
		Type:     myUserModel.Type,
	}
	return user
}
