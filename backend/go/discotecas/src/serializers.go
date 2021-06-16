package discotecas

// import (
// 	"github.com/gin-gonic/gin"
// 	"goApp/common"
// )


// // Declare your response schema here
// type Discotecas struct {
// 	Id          uint
// 	Name        string   `json:"name"`
// 	Company     string   `json:"company"`
// 	Views		int	     `json:"views"`
// }

// type DiscotecassSerializer struct {
// 	C        *gin.Context
// 	Discotecas []Discotecas
// }

// type DiscotecasSerializer struct {
// 	C *gin.Context
// 	Discotecas
// }


// func (s *DiscotecasSerializer) Response() Discotecas {
// 	myUserModel := s.C.MustGet("my_user_model").(UserModel)
// 	response := Discotecas {
// 		Id:          s.Id,
// 		Name:        s.Name
// 		Company:     s.Company
// 		// Favorite:    s.isFavoriteBy(GetDiscotecasUserModel(myUserModel)),
// 		Views:	     s.Views,
// 	}
	
// 	return response
// }


// func (s *DiscotecassSerializer) Response() []Discotecas {
// 	response := []Discotecas{}
// 	for _, discotecas := range s.Discotecass {
// 		serializer := DiscotecasSerializer{s.C, discotecas}
// 		response = append(response, serializer.Response())
// 	}
// 	return response
// }
