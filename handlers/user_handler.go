package handlers 

import(
	"log"
	"net/http"
	"go-sqlite/services"
	"go-sqlite/models"
	"encoding/json"
	
)

type UserHandler struct{
	service *services.UserServices
}

func NewUserHandler(service *services.UserServices) *UserHandler{
	return &UserHandler{service:service}
}

func (handler *UserHandler) InsertUser(writer http.ResponseWriter, request *http.Request){
	
	if request.Method!=http.MethodPost{
			http.Error(writer,"Invalid Method type",405)
			log.Println("Invalid Method type")
			return 
		}

	var newuser models.Users

	err :=json.NewDecoder(request.Body).Decode(&newuser)
	if err != nil {
			http.Error(writer,"Invalid body or empty body",400)
			log.Println("error in fetching the data")
			return
		}
	err = handler.service.InsertUser(newuser)	
	if err!=nil{
		log.Println("error in service function calling ")
		return
	}
	json.NewEncoder(writer).Encode(map[string]interface{}{
			"message":  "the user inserted succesfully into database ",
			
		})
}