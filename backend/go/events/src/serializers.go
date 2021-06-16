package events

// import (
// 	"github.com/gin-gonic/gin"
// 	"goApp/common"
// )


// // Declare your response schema here
// type Events struct {
// 	Id          uint
// 	Name        string   `json:"name"`
// 	Company     string   `json:"company"`
// 	Views		int	     `json:"views"`
// }

// type EventssSerializer struct {
// 	C        *gin.Context
// 	Events []Events
// }

// type EventsSerializer struct {
// 	C *gin.Context
// 	Events
// }


// func (s *EventsSerializer) Response() Events {
// 	myUser := s.C.MustGet("my_user_model").(User)
// 	response := Events {
// 		Id:          s.Id,
// 		Name:        s.Name
// 		Company:     s.Company
// 		// Favorite:    s.isFavoriteBy(GetEventsUser(myUser)),
// 		Views:	     s.Views,
// 	}
	
// 	return response
// }


// func (s *EventssSerializer) Response() []Events {
// 	response := []Events{}
// 	for _, events := range s.Eventss {
// 		serializer := EventsSerializer{s.C, events}
// 		response = append(response, serializer.Response())
// 	}
// 	return response
// }
